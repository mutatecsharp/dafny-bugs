// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, p1, p2, p3, globalState) {
      if (_dafny.Seq.contains(_dafny.Seq.of(false, !(true)), true)) {
        return (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-675))).cardinality())).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("g")).length), new BigNumber(873), new BigNumber(829), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),true)).length))).length));
      } else {
        return ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(431))).cardinality()))).multipliedBy(new BigNumber(590));
      }
    };
    static fm1(p0, globalState) {
      return !(false);
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(914))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-832)), function (_0_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false)).length)))));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1((_dafny.Set.fromElements(true, false)).IsSubsetOf(_dafny.Set.fromElements(true)));
    };
    static fm6(globalState) {
      return _dafny.Set.fromElements(new BigNumber(((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(716), new BigNumber(849))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(716)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(849)))) {
            _coll0.push([_module.__default.safeDivisionInt(_1_v0, new BigNumber((_dafny.Seq.UnicodeFromString("ja")).length)),new BigNumber(-249)]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(125),new BigNumber((_dafny.Seq.UnicodeFromString("m")).length)))).length));
    };
    static fm7(globalState) {
      return (_dafny.MultiSet.fromElements(true)).Union((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false), true))).Difference(_dafny.MultiSet.fromElements(true)));
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("khrapsoyo"), _dafny.Seq.UnicodeFromString("l"));
    };
    static fm9(p0, globalState) {
      return new _dafny.CodePoint('r'.codePointAt(0));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(410), new BigNumber(725))) {
          let _2_v0 = _compr_1;
          if (((new BigNumber(410)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(725)))) {
            _coll1.add((_2_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tsnmoy")).length)));
          }
        }
        return _coll1;
      }()).IsSubsetOf(function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(291), new BigNumber(256))) {
            let _3_v1 = _compr_3;
            if (((new BigNumber(291)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(256)))) {
              _coll3.push([_module.__default.safeDivisionInt(_3_v1, new BigNumber((_dafny.Set.fromElements(new BigNumber(-295), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-593), new BigNumber(360))).cardinality())))).length)),true]);
            }
          }
          return _coll3;
        }()).Keys.Elements) {
          let _4_v2 = _compr_2;
          if ((function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of _dafny.IntegerRange(new BigNumber(291), new BigNumber(256))) {
              let _3_v1 = _compr_4;
              if (((new BigNumber(291)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(256)))) {
                _coll4.push([_module.__default.safeDivisionInt(_3_v1, new BigNumber((_dafny.Set.fromElements(new BigNumber(-295), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-593), new BigNumber(360))).cardinality())))).length)),true]);
              }
            }
            return _coll4;
          }()).contains(_4_v2)) {
            _coll2.add((_4_v2).multipliedBy(new BigNumber(208)));
          }
        }
        return _coll2;
      }()), false, !(_dafny.Set.fromElements(new BigNumber(-756), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-709)), function (_5_i0) {
        return new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("rdw")).length),new _dafny.CodePoint('m'.codePointAt(0)))).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(497)), function (_6_i1) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("smixm")).length);
        })).length), new BigNumber(338))).length);
      })).length))).contains(new BigNumber(478)), true);
    };
    static fm13(globalState) {
      return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,!(false)), _dafny.Map.Empty.slice().updateUnsafe(true,false))).Union((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Difference(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Elements) {
            let _7_v0 = _compr_6;
            if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false))).contains(_7_v0)) {
              _coll6.push([_7_v0,true]);
            }
          }
          return _coll6;
        }()).Keys.Elements) {
          let _8_v1 = _compr_5;
          if ((function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Elements) {
              let _7_v0 = _compr_7;
              if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false))).contains(_7_v0)) {
                _coll7.push([_7_v0,true]);
              }
            }
            return _coll7;
          }()).contains(_8_v1)) {
            _coll5.add(_8_v1);
          }
        }
        return _coll5;
      }()));
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dybqgrbwg"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wxkb"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(756)), function (_9_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })));
    };
    static fm15(p0, p1, globalState) {
      return _dafny.Seq.Concat((_dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), function (_10_i0) {
        return new BigNumber(-520);
      })), _dafny.Seq.of(new BigNumber(421), new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(469)));
    };
    static fm16(globalState) {
      let _source0 = _module.D13.create_DC38(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of _dafny.IntegerRange(new BigNumber(874), new BigNumber(620))) {
    let _11_v0 = _compr_8;
    if (((new BigNumber(874)).isLessThanOrEqualTo(_11_v0)) && ((_11_v0).isLessThan(new BigNumber(620)))) {
      _coll8.push([(_11_v0).multipliedBy(new BigNumber(751)),new BigNumber(-828)]);
    }
  }
  return _coll8;
}()).length)));
      if (_source0.is_DC39) {
        if (false) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }
      } else if (_source0.is_DC40) {
        let _12___mcc_h0 = (_source0).cf56;
        let _13___mcc_h1 = (_source0).cf57;
        let _14___mcc_h2 = (_source0).cf58;
        let _15_cf58 = _14___mcc_h2;
        let _16_cf57 = _13___mcc_h1;
        let _17_cf56 = _12___mcc_h0;
        return new _dafny.CodePoint('p'.codePointAt(0));
      } else {
        let _18___mcc_h3 = (_source0).cf55;
        let _19_cf55 = _18___mcc_h3;
        return new _dafny.CodePoint('u'.codePointAt(0));
      }
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(true))).Difference((_dafny.Set.fromElements(false)).Union(_dafny.Set.fromElements(true)));
    };
    static fm19(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((false) ? (new BigNumber(822)) : (new BigNumber(-43))),new _dafny.CodePoint('t'.codePointAt(0)));
    };
    static fm20(p0, p1, p2, globalState) {
      if (!_dafny.areEqual(_dafny.Seq.of(_dafny.Set.fromElements(!(false)), _dafny.Set.fromElements(false), _dafny.Set.fromElements(true), _dafny.Set.fromElements(!(false))), _dafny.Seq.of(_dafny.Set.fromElements(true, !(!(true)))))) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }
    };
    static fm21(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(791)), function (_20_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("curdiyset")), _dafny.Seq.UnicodeFromString("o"));
    };
    static fm22(p0, globalState) {
      let _source1 = _module.D12.create_DC37(_module.D12.create_DC36(true, new BigNumber(-997), false));
      if (_source1.is_DC34) {
        let _21___mcc_h0 = (_source1).cf48;
        let _22___mcc_h1 = (_source1).cf49;
        let _23___mcc_h2 = (_source1).cf50;
        let _24_cf50 = _23___mcc_h2;
        let _25_cf49 = _22___mcc_h1;
        let _26_cf48 = _21___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(26)), function (_27_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })).length), new BigNumber(192))).length)),true);
      } else if (_source1.is_DC35) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(882),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(144),true));
      } else if (_source1.is_DC36) {
        let _28___mcc_h3 = (_source1).cf51;
        let _29___mcc_h4 = (_source1).cf52;
        let _30___mcc_h5 = (_source1).cf53;
        let _31_cf53 = _30___mcc_h5;
        let _32_cf52 = _29___mcc_h4;
        let _33_cf51 = _28___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(_32_cf52,_33_cf51);
      } else if (_source1.is_DC33) {
        let _34___mcc_h6 = (_source1).cf47;
        let _35_cf47 = _34___mcc_h6;
        return (function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(997), new BigNumber(164))) {
            let _36_v0 = _compr_9;
            if (((new BigNumber(997)).isLessThanOrEqualTo(_36_v0)) && ((_36_v0).isLessThan(new BigNumber(164)))) {
              _coll9.push([(_36_v0).minus(new BigNumber(492)),false]);
            }
          }
          return _coll9;
        }()).Merge(function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(766), new BigNumber(670))) {
            let _37_v1 = _compr_10;
            if (((new BigNumber(766)).isLessThanOrEqualTo(_37_v1)) && ((_37_v1).isLessThan(new BigNumber(670)))) {
              _coll10.push([(_37_v1).multipliedBy(new BigNumber((_35_cf47).length)),!(!(true))]);
            }
          }
          return _coll10;
        }());
      } else {
        let _38___mcc_h7 = (_source1).cf54;
        let _39_cf54 = _38___mcc_h7;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("slrpor")).length),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(765),false));
      }
    };
    static fm23(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("qgn")).length))).Difference((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-41))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber(458))).Difference(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false))).length)),new BigNumber((_dafny.Set.fromElements(false, false, false, false)).length)),true)).Keys.Elements) {
          let _40_v0 = _compr_11;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false))).length)),new BigNumber((_dafny.Set.fromElements(false, false, false, false)).length)),true)).contains(_40_v0)) {
            _coll11.add(_40_v0);
          }
        }
        return _coll11;
      }()).length))));
    };
    static fm24(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(!(false)) || ((_module.D9.create_DC24(!(!(true)), false, true)).dtor_cf35),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(855)));
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.Concat(_dafny.Seq.of(false, true), _dafny.Seq.of((_module.D12.create_DC34(_module.D0.create_DC2(_module.D0.create_DC1(false)), false, true)).dtor_cf50, true, !(true), false)));
    };
    static fm26(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D3.create_DC10((_module.D3.create_DC10(_module.D3.create_DC9(new BigNumber(528)))).dtor_cf12);
      if (_source2.is_DC9) {
        let _41___mcc_h0 = (_source2).cf11;
        let _42_cf11 = _41___mcc_h0;
        return _module.D2.create_DC7(new BigNumber(-394), function () {
  let _coll12 = new _dafny.Map();
  for (const _compr_12 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(true),true))).Elements) {
    let _43_v0 = _compr_12;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(true),true)), _43_v0)) {
      _coll12.push([_43_v0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_42_cf11,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true, true, true)).length)))).length)]);
    }
  }
  return _coll12;
}(), new BigNumber(890));
      } else if (_source2.is_DC8) {
        let _44___mcc_h1 = (_source2).cf10;
        let _45_cf10 = _44___mcc_h1;
        return _module.D2.create_DC7(new BigNumber(338), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,false),new BigNumber(-77)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(373)), function (_46_i0) {
  return new _dafny.CodePoint('o'.codePointAt(0));
})).length));
      } else {
        let _47___mcc_h2 = (_source2).cf12;
        let _48_cf12 = _47___mcc_h2;
        return _module.D2.create_DC7(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(115)), function (_49_i1) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})).length), function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,true))).Elements) {
    let _50_v1 = _compr_13;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,true)), _50_v1)) {
      _coll13.push([_50_v1,new BigNumber(762)]);
    }
  }
  return _coll13;
}(), new BigNumber(107));
      }
    };
    static fm27(globalState) {
      if (!(false) || (!(false))) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("j"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(455)), function (_51_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }));
      } else {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(838)), function (_52_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        });
      }
    };
    static fm28(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)));
    };
    static fm29(globalState) {
      if (!(!((!(false)) && (true)))) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }
    };
    static fm30(globalState) {
      let _source3 = ((false) ? (_module.D11.create_DC30()) : (_module.D11.create_DC30()));
      if (_source3.is_DC30) {
        return _dafny.Set.fromElements(new BigNumber(-94), new BigNumber(271));
      } else if (_source3.is_DC31) {
        let _53___mcc_h0 = (_source3).cf44;
        let _54___mcc_h1 = (_source3).cf45;
        let _55_cf45 = _54___mcc_h1;
        let _56_cf44 = _53___mcc_h0;
        return (_dafny.Set.fromElements(_56_cf44)).Intersect(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length))));
      } else if (_source3.is_DC29) {
        let _57___mcc_h2 = (_source3).cf43;
        let _58_cf43 = _57___mcc_h2;
        return (_dafny.Set.fromElements(new BigNumber(413))).Union(_dafny.Set.fromElements(new BigNumber(-448), new BigNumber((function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(591), new BigNumber(801))) {
            let _59_v0 = _compr_14;
            if (((new BigNumber(591)).isLessThanOrEqualTo(_59_v0)) && ((_59_v0).isLessThan(new BigNumber(801)))) {
              _coll14.push([(_59_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("aydffd")).length)),new BigNumber((_dafny.Seq.UnicodeFromString("ikjyjyifl")).length)]);
            }
          }
          return _coll14;
        }()).length)));
      } else {
        let _60___mcc_h3 = (_source3).cf46;
        let _61_cf46 = _60___mcc_h3;
        return (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ba")).length))).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false, false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ryykijq")).length),new _dafny.CodePoint('t'.codePointAt(0)))).length)));
      }
    };
    static fm31(globalState) {
      return ((_dafny.MultiSet.fromElements(_module.D0.create_DC2(_module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-475)))).cardinality()),false))), _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jgwt")).length),true)))), _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC0(function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-981)), function (_62_i0) {
    return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), function (_63_i1) {
      return new _dafny.CodePoint('c'.codePointAt(0));
    })).length))).cardinality())),!(false))).length);
  })).Elements) {
    let _64_v0 = _compr_15;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-981)), function (_62_i0) {
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), function (_63_i1) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length))).cardinality())),!(false))).length);
    }), _64_v0)) {
      _coll15.push([(_64_v0).plus(new BigNumber(806)),true]);
    }
  }
  return _coll15;
}()))))).Intersect(_dafny.MultiSet.fromElements(_module.D0.create_DC2(_module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)),false)))))).Intersect((_dafny.MultiSet.fromElements(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1(true))), _module.D0.create_DC2(_module.D0.create_DC1(false)), _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1(true))))).Intersect(_dafny.MultiSet.fromElements(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1(true))))))));
    };
    static fm32(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D3.create_DC9((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), function (_65_i0) {
  return new _dafny.CodePoint('a'.codePointAt(0));
})).length),false)).length))).length))), _module.D3.create_DC9((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber(678), new BigNumber(813), new BigNumber((_dafny.Set.fromElements(new BigNumber(973), new BigNumber(62), (_dafny.ZERO).minus(new BigNumber(-221)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length), new BigNumber((_dafny.Seq.of(true)).length))).length))).cardinality())))), _dafny.Seq.of(_module.D3.create_DC9(new BigNumber(-407))));
    };
    static fm33(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("pptit")).length),_module.D8.create_DC21(false, false, _dafny.Seq.of(false, true), true, new BigNumber(249)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-630),new BigNumber(-844))).length),_module.D8.create_DC21(true, true, _dafny.Seq.of(false), !(false), new BigNumber(-398))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(123),_module.D8.create_DC21(true, true, _dafny.Seq.of(false), false, new BigNumber((_dafny.Seq.of(new BigNumber(48))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ska")).length),_module.D8.create_DC21(true, !(!(false)), _dafny.Seq.of(false, true, false, false, true), true, new BigNumber(-541)))));
    };
    static fm34(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false, true))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(!(true), !(false), false, false, !(true)),new BigNumber(299)));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _66_v0;
      _66_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.fromElements(p0, p0, p0, false, p0));
      let _67_v1;
      _67_v1 = _module.D0.create_DC1(p0);
      let _68_v2;
      _68_v2 = _dafny.MultiSet.fromElements(p0, (_67_v1).dtor_cf1);
      let _69_v3;
      _69_v3 = _module.D10.create_DC27(p0, new BigNumber(921), p1, (((_66_v0).contains(!(!(true)))) ? ((_66_v0).get(!(!(true)))) : (_68_v2)));
      let _70_v4;
      _70_v4 = _dafny.Seq.of(p0);
      let _71_v5;
      _71_v5 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_70_v4, _module.__default.safeIndex(p2, new BigNumber((_70_v4).length)), !(p0))).length)), p2, p2, new BigNumber(947));
      if ((_71_v5).contains(_module.__default.safeModuloInt((_dafny.ZERO).minus((((p3).contains(true)) ? ((p3).get(true)) : (new BigNumber(((_69_v3).dtor_cf41).cardinality())))), p2))) {
        (globalState).f0 = p0;
        let _72_v6;
        let _nw0 = new _module.C2();
        _nw0.__ctor(p0);
        _72_v6 = _nw0;
        _72_v6 = _72_v6;
        let _73_v7;
        _73_v7 = _dafny.Seq.UnicodeFromString("osmkakap");
        _73_v7 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), ((_74_p1) => function (_75_i0) {
          return _74_p1;
        })(p1)), _dafny.Seq.UnicodeFromString("baefwmqnu"));
        let _76_v8;
        _76_v8 = _dafny.MultiSet.fromElements(p1, p1);
        let _77_v9;
        let _init0 = ((_78_v6) => function (_79_i1) {
          return !((_78_v6).f13);
        })(_72_v6);
        let _nw1 = Array((new BigNumber(24)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _77_v9 = _nw1;
        let _80_v10;
        let _nw2 = new _module.C0();
        _nw2.__ctor(!(p2).isEqualTo(p2), (_module.D4.create_DC13(new BigNumber((_76_v8).cardinality()), _77_v9)).dtor_cf20);
        _80_v10 = _nw2;
        let _81_v11;
        let _nw3 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _81_v11 = _nw3;
        let _index0 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_81_v11).length));
        (_81_v11)[_index0] = _module.__default.fm14(_73_v7, p2, globalState);
        let _index1 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_81_v11).length));
        (_81_v11)[_index1] = _dafny.Seq.Concat(_73_v7, _73_v7);
      } else {
        let _82_v12;
        _82_v12 = _dafny.Map.Empty.slice().updateUnsafe(p2,!((_69_v3).dtor_cf38) || (p0));
        _82_v12 = _82_v12;
        let _83_v13;
        _83_v13 = _dafny.Seq.of(p2);
        let _84_v14;
        _84_v14 = _dafny.Seq.UnicodeFromString("yxwm");
        let _rhs0 = _83_v13;
        let _rhs1 = !(p0);
        let _rhs2 = p0;
        let _rhs3 = new BigNumber((_84_v14).length);
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        let _lhs2 = globalState;
        _83_v13 = _rhs0;
        _lhs0.f0 = _rhs1;
        _lhs1.f0 = _rhs2;
        _lhs2.f5 = _rhs3;
        let _85_v15;
        _85_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm0(new BigNumber((_84_v14).length), p2, new BigNumber((_84_v14).length), p0, globalState)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), ((_86_p2) => function (_87_i2) {
          return _86_p2;
        })(p2))).length));
        _85_v15 = (_85_v15).update(new BigNumber((_dafny.Set.fromElements(false, p0)).length), new BigNumber((_82_v12).length));
        let _88_v16;
        let _init1 = ((_89_p0) => function (_90_i3) {
          return _89_p0;
        })(p0);
        let _nw4 = Array((new BigNumber(23)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw4.length); _i0_1++) {
          _nw4[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _88_v16 = _nw4;
        let _index2 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_88_v16).length));
        (_88_v16)[_index2] = p0;
        let _index3 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_88_v16).length));
        (_88_v16)[_index3] = (p2).isLessThan(p2);
        let _index4 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_88_v16).length));
        (_88_v16)[_index4] = !_dafny.Seq.contains(_70_v4, p0);
      }
      (globalState).f2 = p0;
      if (_module.__default.fm1(p0, globalState)) {
        let _91_v17;
        let _nw5 = Array((new BigNumber(14)).toNumber()).fill(false);
        _91_v17 = _nw5;
        let _index5 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length));
        (_91_v17)[_index5] = false;
        let _index6 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length));
        (_91_v17)[_index6] = !(false);
        if (p0) {
          let _92_v18;
          let _init2 = ((_93_p2) => function (_94_i4) {
            return _module.__default.safeModuloInt(_94_i4, _93_p2);
          })(p2);
          let _nw6 = Array((new BigNumber(5)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
            _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _92_v18 = _nw6;
          let _95_v19;
          _95_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,_92_v18);
          let _96_v20;
          _96_v20 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
          let _97_v21;
          _97_v21 = _dafny.Seq.UnicodeFromString("ikfpgwiid");
          let _98_v22;
          let _nw7 = new _module.C3();
          _nw7.__ctor(p2);
          _98_v22 = _nw7;
          let _99_v23;
          _99_v23 = _dafny.Map.Empty.slice().updateUnsafe(_98_v22,p2);
          let _100_v24;
          let _nw8 = new _module.C2();
          _nw8.__ctor(p0);
          _100_v24 = _nw8;
          let _101_v25;
          _101_v25 = _dafny.Map.Empty.slice().updateUnsafe(p2,_100_v24);
          let _102_v26;
          _102_v26 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber((_101_v25).length));
          let _103_v27;
          _103_v27 = _dafny.Map.Empty.slice().updateUnsafe(_91_v17,new BigNumber(806));
          let _104_v28;
          _104_v28 = _dafny.Set.fromElements((_91_v17)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length))], p0);
          let _105_v29;
          let _nw9 = Array((new BigNumber(21)).toNumber());
          _nw9[(_dafny.ZERO).toNumber()] = p2;
          _nw9[(_dafny.ONE).toNumber()] = p2;
          _nw9[(new BigNumber(2)).toNumber()] = (p2).minus(new BigNumber((_95_v19).length));
          _nw9[(new BigNumber(3)).toNumber()] = p2;
          _nw9[(new BigNumber(4)).toNumber()] = p2;
          _nw9[(new BigNumber(5)).toNumber()] = p2;
          _nw9[(new BigNumber(6)).toNumber()] = p2;
          _nw9[(new BigNumber(7)).toNumber()] = p2;
          _nw9[(new BigNumber(8)).toNumber()] = (((p3).contains(p0)) ? ((p3).get(p0)) : ((_dafny.ZERO).minus(_module.__default.fm0(new BigNumber((_96_v20).length), p2, new BigNumber((_97_v21).length), p0, globalState))));
          _nw9[(new BigNumber(9)).toNumber()] = new BigNumber((_99_v23).length);
          _nw9[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw9[(new BigNumber(11)).toNumber()] = p2;
          _nw9[(new BigNumber(12)).toNumber()] = p2;
          _nw9[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(p2, (_dafny.ZERO).minus(p2));
          _nw9[(new BigNumber(14)).toNumber()] = (((_102_v26).contains(p1)) ? ((_102_v26).get(p1)) : (p2));
          _nw9[(new BigNumber(15)).toNumber()] = p2;
          _nw9[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw9[(new BigNumber(17)).toNumber()] = p2;
          _nw9[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_103_v27).Merge(_103_v27)).length));
          _nw9[(new BigNumber(19)).toNumber()] = new BigNumber((_104_v28).length);
          _nw9[(new BigNumber(20)).toNumber()] = p2;
          _105_v29 = _nw9;
          let _106_v30;
          let _nw10 = new _module.C0();
          _nw10.__ctor((_91_v17)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length))], _91_v17);
          _106_v30 = _nw10;
          let _107_v31;
          _107_v31 = _dafny.Set.fromElements(_106_v30);
          let _108_v32;
          _108_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(p2, new BigNumber((_107_v31).length))).cardinality()),_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_109_i5) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("pbcudgtk")).length);
          }));
          let _110_v33;
          _110_v33 = _dafny.Seq.of(p2);
          let _index7 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length));
          (_105_v29)[_index7] = new BigNumber((_dafny.Seq.Concat((((_108_v32).contains(p2)) ? ((_108_v32).get(p2)) : (_dafny.Seq.of(new BigNumber(552)))), _110_v33)).length);
          let _111_v34;
          let _init3 = ((_112_p1) => function (_113_i6) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), ((_114_p1) => function (_115_i7) {
              return _114_p1;
            })(_112_p1));
          })(p1);
          let _nw11 = Array((new BigNumber(25)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw11.length); _i0_3++) {
            _nw11[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _111_v34 = _nw11;
          let _index8 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length));
          let _rhs4 = p2;
          let _rhs5 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("agjcjgg"), _dafny.Seq.Concat(_97_v21, _dafny.Seq.Create(_module.__default.abs(new BigNumber(944)), ((_116_p1) => function (_117_i8) {
            return _116_p1;
          })(p1))));
          let _rhs6 = p2;
          let _rhs7 = _111_v34;
          let _lhs3 = _105_v29;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length));
          let _lhs5 = globalState;
          let _lhs6 = globalState;
          _lhs3[_lhs4] = _rhs4;
          _lhs5.f2 = _rhs5;
          _lhs6.f5 = _rhs6;
          _111_v34 = _rhs7;
          let _index9 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length));
          let _rhs8 = p0;
          let _rhs9 = _module.__default.fm0((_105_v29)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length))], (_105_v29)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length))], p2, p0, globalState);
          let _rhs10 = ((_105_v29)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length))]).multipliedBy((_105_v29)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length))]);
          let _lhs7 = globalState;
          let _lhs8 = globalState;
          let _lhs9 = _105_v29;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length));
          _lhs7.f2 = _rhs8;
          _lhs8.f5 = _rhs9;
          _lhs9[_lhs10] = _rhs10;
          let _118_v35;
          let _nw12 = new _module.C4();
          _nw12.__ctor((_105_v29)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_105_v29).length))]);
          _118_v35 = _nw12;
          let _119_v36;
          _119_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_118_v35);
          let _120_v37;
          _120_v37 = _dafny.Seq.of(_118_v35, (((_119_v36).contains(new BigNumber((_104_v28).length))) ? ((_119_v36).get(new BigNumber((_104_v28).length))) : (_118_v35)), _118_v35);
          _120_v37 = _dafny.Seq.Concat(_120_v37, _dafny.Seq.update(_120_v37, _module.__default.safeIndex(p2, new BigNumber((_120_v37).length)), _118_v35));
          let _121_v38;
          _121_v38 = new _dafny.CodePoint('s'.codePointAt(0));
          _121_v38 = (((_100_v24).f13) ? (p1) : (p1));
          (globalState).f0 = !(p3).contains(true);
        } else {
          let _122_v39;
          let _nw13 = Array((new BigNumber(11)).toNumber());
          _122_v39 = _nw13;
          let _123_v40;
          _123_v40 = _dafny.Seq.of(p2);
          let _124_v41;
          let _nw14 = new _module.C2();
          _nw14.__ctor((_70_v4)[_module.__default.safeIndex((_123_v40)[_module.__default.safeIndex(p2, new BigNumber((_123_v40).length))], new BigNumber((_70_v4).length))]);
          _124_v41 = _nw14;
          let _index10 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_122_v39).length));
          (_122_v39)[_index10] = _124_v41;
          let _index11 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_122_v39).length));
          let _nw15 = new _module.C2();
          _nw15.__ctor((_124_v41).f13);
          (_122_v39)[_index11] = _nw15;
          (globalState).f2 = p0;
          let _125_v42;
          let _nw16 = new _module.C2();
          _nw16.__ctor(p0);
          _125_v42 = _nw16;
          let _126_v43;
          _126_v43 = _dafny.Map.Empty.slice().updateUnsafe(_125_v42,p2);
          let _127_v44;
          _127_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_126_v43).length),(_91_v17)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length))]);
          let _128_v45;
          _128_v45 = _module.D0.create_DC0(_127_v44);
          _68_v2 = _dafny.MultiSet.fromElements(!((false) && ((_91_v17)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length))])), !((_124_v41).fm4(_128_v45, _128_v45, _68_v2, globalState)));
          let _129_v46;
          _129_v46 = _dafny.Seq.UnicodeFromString("colybgq");
          let _130_v47;
          _130_v47 = _dafny.Map.Empty.slice().updateUnsafe(_129_v46,_dafny.Seq.Concat(_123_v40, _dafny.Seq.of(new BigNumber(652))));
          let _131_v48;
          _131_v48 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(131)), ((_132_p2) => function (_133_i9) {
            return _132_p2;
          })(p2)));
          _130_v47 = (_130_v47).update(_dafny.Seq.UnicodeFromString("tqy"), (_131_v48)[_module.__default.safeIndex(new BigNumber(-912), new BigNumber((_131_v48).length))]);
          let _134_v49;
          let _nw17 = new _module.C3();
          _nw17.__ctor(p2);
          _134_v49 = _nw17;
        }
        let _135_v50;
        let _nw18 = new _module.C1();
        _nw18.__ctor();
        _135_v50 = _nw18;
        (globalState).f0 = (_68_v2).IsSubsetOf(_68_v2);
        let _136_v51;
        _136_v51 = _module.D7.create_DC18();
        let _137_v52;
        _137_v52 = _dafny.Seq.of(_136_v51);
        let _138_v53;
        _138_v53 = _dafny.Seq.of(_137_v52);
        let _139_v54;
        _139_v54 = _dafny.Map.Empty.slice().updateUnsafe((_91_v17)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length))],_138_v53);
        let _index12 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_91_v17).length));
        (_91_v17)[_index12] = (new BigNumber(((((_139_v54).contains(p0)) ? ((_139_v54).get(p0)) : (_138_v53))).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(((true) ? (new BigNumber(-279)) : (p2))));
      } else {
        let _140_v55;
        _140_v55 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _141_v56;
        let _nw19 = new _module.C4();
        _nw19.__ctor(new BigNumber((_140_v55).length));
        _141_v56 = _nw19;
        let _142_v57;
        let _nw20 = Array((new BigNumber(15)).toNumber());
        _nw20[(_dafny.ZERO).toNumber()] = _141_v56;
        _nw20[(_dafny.ONE).toNumber()] = _141_v56;
        _nw20[(new BigNumber(2)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(3)).toNumber()] = ((!(p0)) ? (_141_v56) : (_141_v56));
        _nw20[(new BigNumber(4)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(5)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(6)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(7)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(8)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(9)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(10)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(11)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(12)).toNumber()] = _141_v56;
        _nw20[(new BigNumber(13)).toNumber()] = ((p0) ? (_141_v56) : (_141_v56));
        _nw20[(new BigNumber(14)).toNumber()] = _141_v56;
        _142_v57 = _nw20;
        let _index13 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_142_v57).length));
        (_142_v57)[_index13] = _141_v56;
        let _143_v58;
        _143_v58 = _dafny.Seq.UnicodeFromString("tlsst");
        let _144_v59;
        _144_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_143_v58).length));
        let _145_v60;
        let _nw21 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _145_v60 = _nw21;
        let _index14 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_145_v60).length));
        (_145_v60)[_index14] = (_dafny.ZERO).minus((_141_v56).f9);
        let _146_v61;
        let _init4 = ((_147_p0) => function (_148_i10) {
          return _147_p0;
        })(p0);
        let _nw22 = Array((new BigNumber(23)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw22.length); _i0_4++) {
          _nw22[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _146_v61 = _nw22;
        let _149_v62;
        _149_v62 = _module.D4.create_DC13(p2, _146_v61);
        let _150_v63;
        _150_v63 = _141_v56;
        let _151_v64;
        _151_v64 = (_150_v63);
        let _152_v65;
        _152_v65 = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('y'.codePointAt(0)));
        let _153_v66;
        _153_v66 = _module.D2.create_DC6(p0);
        let _154_v67;
        let _nw23 = new _module.C2();
        _nw23.__ctor(p0);
        _154_v67 = _nw23;
        let _155_v68;
        _155_v68 = _dafny.Set.fromElements(_154_v67, _154_v67, _154_v67);
        let _index15 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_142_v57).length));
        let _index16 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_145_v60).length));
        let _rhs11 = (_150_v63);
        let _rhs12 = _144_v59;
        let _rhs13 = new BigNumber((_module.__default.fm8(p0, _152_v65, (_153_v66).dtor_cf6, globalState)).length);
        let _rhs14 = _149_v62;
        let _rhs15 = ((_155_v68).IsDisjointFrom(_155_v68)) && ((_154_v67).f13);
        let _lhs11 = _142_v57;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_142_v57).length));
        let _lhs13 = _145_v60;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_145_v60).length));
        let _lhs15 = globalState;
        _lhs11[_lhs12] = _rhs11;
        _144_v59 = _rhs12;
        _lhs13[_lhs14] = _rhs13;
        _149_v62 = _rhs14;
        _lhs15.f2 = _rhs15;
        let _156_v69;
        let _nw24 = new _module.C0();
        _nw24.__ctor(p0, _146_v61);
        _156_v69 = _nw24;
        let _157_v70;
        let _nw25 = new _module.C1();
        _nw25.__ctor();
        _157_v70 = _nw25;
        let _158_v71;
        _158_v71 = _dafny.Set.fromElements(_157_v70, _157_v70, _157_v70, _157_v70);
        let _159_v72;
        _159_v72 = _dafny.Seq.of(new BigNumber((_158_v71).length), p2);
        let _index17 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_145_v60).length));
        (_145_v60)[_index17] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeModuloInt((_141_v56).f9, p2), new BigNumber((_159_v72).length)));
        _140_v55 = _140_v55;
        let _160_v74;
        _160_v74 = _dafny.Seq.of(_dafny.Seq.of((_156_v69).f10, (_154_v67).f13));
        let _161_v75;
        _161_v75 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of (_160_v74).Elements) {
            let _162_v73 = _compr_16;
            if (_dafny.Seq.contains(_160_v74, _162_v73)) {
              _coll16.push([_162_v73,new BigNumber((_159_v72).length)]);
            }
          }
          return _coll16;
        }(),_145_v60);
        let _163_v76;
        _163_v76 = _dafny.Map.Empty.slice().updateUnsafe(_70_v4,(_145_v60)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_145_v60).length))]);
        _145_v60 = (((_161_v75).contains((_163_v76).Merge(_163_v76))) ? ((_161_v75).get((_163_v76).Merge(_163_v76))) : (_145_v60));
      }
      let _164_v77;
      let _init5 = ((_165_p0) => function (_166_i11) {
        return _module.__default.safeDivisionInt(_166_i11, new BigNumber((_dafny.Set.fromElements(_165_p0)).length));
      })(p0);
      let _nw26 = Array((new BigNumber(20)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw26.length); _i0_5++) {
        _nw26[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _164_v77 = _nw26;
      let _167_v78;
      _167_v78 = _dafny.Map.Empty.slice().updateUnsafe(p0,_164_v77);
      (globalState).f5 = _module.__default.safeModuloInt((p2).multipliedBy(new BigNumber((_167_v78).length)), _module.__default.fm0(p2, _module.__default.fm0(new BigNumber(818), p2, new BigNumber((_dafny.Seq.UnicodeFromString("dmpivcli")).length), p0, globalState), p2, p0, globalState));
      let _168_v79;
      let _nw27 = new _module.C4();
      _nw27.__ctor((p2).minus(p2));
      _168_v79 = _nw27;
      (globalState).f0 = false;
      return;
    }
    static Main(__noArgsParameter) {
      let _169_v0;
      _169_v0 = new BigNumber(292);
      let _170_v1;
      _170_v1 = _dafny.Seq.of(_169_v0, _169_v0);
      let _171_v2;
      _171_v2 = false;
      let _172_v3;
      let _nw28 = Array((new BigNumber(3)).toNumber());
      _nw28[(_dafny.ZERO).toNumber()] = _169_v0;
      _nw28[(_dafny.ONE).toNumber()] = (_170_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_171_v2,_171_v2)).length), new BigNumber((_170_v1).length))];
      _nw28[(new BigNumber(2)).toNumber()] = _169_v0;
      _172_v3 = _nw28;
      let _173_v4;
      _173_v4 = _dafny.MultiSet.fromElements(_172_v3);
      let _174_v5;
      _174_v5 = _dafny.Seq.of(true);
      let _175_v6;
      _175_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_174_v5).length),_171_v2);
      let _176_v7;
      _176_v7 = _dafny.Map.Empty.slice().updateUnsafe(_169_v0,_dafny.MultiSet.fromElements(_175_v6));
      let _177_v9;
      _177_v9 = _module.D0.create_DC0(function () {
  let _coll17 = new _dafny.Map();
  for (const _compr_17 of _dafny.IntegerRange(new BigNumber(388), new BigNumber(693))) {
    let _178_v8 = _compr_17;
    if (((new BigNumber(388)).isLessThanOrEqualTo(_178_v8)) && ((_178_v8).isLessThan(new BigNumber(693)))) {
      _coll17.push([_module.__default.safeDivisionInt(_178_v8, _169_v0),_171_v2]);
    }
  }
  return _coll17;
}());
      let _179_v10;
      _179_v10 = _dafny.MultiSet.fromElements(_175_v6, (_177_v9).dtor_cf0);
      let _180_v11;
      _180_v11 = _dafny.Set.fromElements(_171_v2, _171_v2);
      let _181_globalState;
      let _nw29 = new _module.GlobalState();
      _nw29.__ctor(true, new BigNumber(-616), false, _173_v4, true, new BigNumber(663), ((((_176_v7).contains(_169_v0)) ? ((_176_v7).get(_169_v0)) : (_179_v10))).Union(_179_v10), false, _180_v11);
      _181_globalState = _nw29;
      let _182_v12;
      _182_v12 = _dafny.Seq.UnicodeFromString("aiwqmuo");
      (_181_globalState).f2 = (_module.__default.safeModuloInt(_169_v0, (_dafny.ZERO).minus(_169_v0))).isLessThan(_module.__default.fm0(new BigNumber((_182_v12).length), _169_v0, _169_v0, _171_v2, _181_globalState));
      let _183_i0;
      _183_i0 = _dafny.ZERO;
      L0: {
        while (_171_v2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_183_i0)) {
              break L0;
            }
            _183_i0 = (_183_i0).plus(_dafny.ONE);
            (_181_globalState).f0 = _module.__default.fm1((_169_v0).isLessThanOrEqualTo(new BigNumber(311)), _181_globalState);
            let _index18 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_172_v3).length));
            (_172_v3)[_index18] = _169_v0;
            let _index19 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_172_v3).length));
            let _rhs16 = _169_v0;
            let _rhs17 = new BigNumber(618);
            let _lhs16 = _172_v3;
            let _lhs17 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_172_v3).length));
            _lhs16[_lhs17] = _rhs16;
            _169_v0 = _rhs17;
            let _184_v13;
            _184_v13 = new _dafny.CodePoint('y'.codePointAt(0));
            let _185_v14;
            _185_v14 = _dafny.Map.Empty.slice().updateUnsafe(_171_v2,new BigNumber((_182_v12).length));
            _module.__default.m0(!((_171_v2) === (_171_v2)), _184_v13, (_dafny.ZERO).minus((((_185_v14).contains(_171_v2)) ? ((_185_v14).get(_171_v2)) : ((_172_v3)[_module.__default.safeIndex(new BigNumber(871), new BigNumber((_172_v3).length))]))), _185_v14, _181_globalState);
            let _186_v15;
            _186_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_169_v0),_170_v1);
            _186_v15 = _186_v15;
          }
        }
      }
      let _187_v17;
      _187_v17 = _dafny.Map.Empty.slice().updateUnsafe(_171_v2,new BigNumber((_182_v12).length));
      _module.__default.m0(true, (_182_v12)[_module.__default.safeIndex(_module.__default.fm0(new BigNumber((function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of (_182_v12).Elements) {
          let _188_v16 = _compr_18;
          if (_dafny.Seq.contains(_182_v12, _188_v16)) {
            _coll18.add(_188_v16);
          }
        }
        return _coll18;
      }()).length), (_dafny.ZERO).minus(_169_v0), _169_v0, _171_v2, _181_globalState), new BigNumber((_182_v12).length))], (_module.__default.fm0(_169_v0, _169_v0, (((_187_v17).contains(_171_v2)) ? ((_187_v17).get(_171_v2)) : (new BigNumber((_182_v12).length))), _171_v2, _181_globalState)).minus(_169_v0), _module.__default.fm2(_171_v2, _171_v2, _171_v2, _187_v17, _181_globalState), _181_globalState);
      let _189_v18;
      _189_v18 = new _dafny.CodePoint('n'.codePointAt(0));
      _module.__default.m0(_171_v2, _189_v18, _module.__default.fm0(_169_v0, _169_v0, _169_v0, _171_v2, _181_globalState), _dafny.Map.Empty.slice().updateUnsafe(_171_v2,_module.__default.fm0(_169_v0, _169_v0, new BigNumber((_182_v12).length), true, _181_globalState)), _181_globalState);
      let _hi0 = new BigNumber(-752);
      for (let _190_i1 = _169_v0; _190_i1.isLessThan(_hi0); _190_i1 = _190_i1.plus(_dafny.ONE)) {
        _169_v0 = _190_i1;
        (_181_globalState).f0 = _171_v2;
        _module.__default.m0(_171_v2, _189_v18, _190_i1, _187_v17, _181_globalState);
        let _191_v19;
        _191_v19 = _dafny.MultiSet.fromElements(_190_i1, _169_v0, _190_i1);
        _191_v19 = _191_v19;
      }
      let _192_v20;
      _192_v20 = _module.D0.create_DC1(_171_v2);
      let _193_v21;
      _193_v21 = _dafny.Map.Empty.slice().updateUnsafe(_171_v2,_192_v20);
      _193_v21 = (_193_v21).update(_171_v2, _module.D0.create_DC1(_171_v2));
      let _source4 = _192_v20;
      if (_source4.is_DC1) {
        let _194___mcc_h0 = (_source4).cf1;
        let _195_cf1 = _194___mcc_h0;
        let _196_v22;
        _196_v22 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("qrvcjj"),_169_v0);
        let _197_v23;
        _197_v23 = _dafny.MultiSet.fromElements(new BigNumber((_196_v22).length));
        let _198_v25;
        _198_v25 = _module.D0.create_DC0(_175_v6);
        let _199_v26;
        _199_v26 = _module.D0.create_DC2(_198_v25);
        let _200_v27;
        _200_v27 = _dafny.Map.Empty.slice().updateUnsafe(_199_v26,_182_v12);
        let _rhs18 = (((_197_v23).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of (_200_v27).Keys.Elements) {
            let _202_v24 = _compr_20;
            if ((_200_v27).contains(_202_v24)) {
              _coll20.push([_202_v24,(_dafny.ZERO).minus(_169_v0)]);
            }
          }
          return _coll20;
        }()).length))))) ? ((_197_v23).get((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of (_200_v27).Keys.Elements) {
            let _201_v24 = _compr_19;
            if ((_200_v27).contains(_201_v24)) {
              _coll19.push([_201_v24,(_dafny.ZERO).minus(_169_v0)]);
            }
          }
          return _coll19;
        }()).length))))) : (_module.__default.fm0(_169_v0, _169_v0, _169_v0, _195_cf1, _181_globalState)));
        let _rhs19 = _module.__default.safeDivisionInt(new BigNumber(550), _module.__default.safeDivisionInt(_169_v0, _169_v0));
        let _lhs18 = _181_globalState;
        _lhs18.f5 = _rhs18;
        _169_v0 = _rhs19;
        let _203_v28;
        _203_v28 = _dafny.Seq.of(_180_v11);
        (_181_globalState).f0 = !(((_180_v11).Intersect((_203_v28)[_module.__default.safeIndex(_169_v0, new BigNumber((_203_v28).length))])).contains((_195_cf1) || (_module.__default.fm1(_195_cf1, _181_globalState))));
        _182_v12 = _182_v12;
        (_181_globalState).f2 = _171_v2;
      } else if (_source4.is_DC0) {
        let _204___mcc_h1 = (_source4).cf0;
        let _205_cf0 = _204___mcc_h1;
        _module.__default.m0(!(_171_v2), _189_v18, _169_v0, _187_v17, _181_globalState);
        _module.__default.m0((!(_171_v2)) === (_171_v2), _189_v18, _169_v0, _187_v17, _181_globalState);
        _182_v12 = _dafny.Seq.Concat(_182_v12, _182_v12);
        (_181_globalState).f5 = _169_v0;
      } else {
        let _206___mcc_h2 = (_source4).cf2;
        let _207_cf2 = _206___mcc_h2;
        _module.__default.m0(_171_v2, _189_v18, (_dafny.ZERO).minus((_169_v0).minus(_169_v0)), _187_v17, _181_globalState);
        let _pat_let_tv0 = _171_v2;
        let _208_v29;
        _208_v29 = _dafny.MultiSet.fromElements(_192_v20, _192_v20, function (_pat_let0_0) {
          return function (_209_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_210_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_210_dt__update_hcf1_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_192_v20));
        let _211_v30;
        _211_v30 = _dafny.Set.fromElements(_172_v3);
        (_181_globalState).f2 = (((_171_v2) ? (_dafny.MultiSet.fromElements(_module.D0.create_DC1((_174_v5)[_module.__default.safeIndex(new BigNumber((_187_v17).length), new BigNumber((_174_v5).length))]))) : (_208_v29))).equals(((false) ? (((_208_v29).update(_192_v20, _module.__default.abs(_169_v0))).update(_module.__default.fm3((_192_v20).dtor_cf1, _169_v0, new BigNumber((_211_v30).length), _169_v0, _181_globalState), _module.__default.abs(_169_v0))) : (_208_v29)));
        _175_v6 = (_175_v6).update(_169_v0, ((_dafny.ZERO).minus(_169_v0)).isLessThan(_169_v0));
        let _212_v31;
        let _nw30 = new _module.C1();
        _nw30.__ctor();
        _212_v31 = _nw30;
      }
      if (false) {
        (_181_globalState).f2 = _171_v2;
        (_181_globalState).f0 = !(_171_v2);
        if (_171_v2) {
          let _213_v32;
          _213_v32 = _dafny.Map.Empty.slice().updateUnsafe(_169_v0,_172_v3);
          let _214_v33;
          _214_v33 = _dafny.MultiSet.fromElements(_171_v2, _module.__default.fm1(_171_v2, _181_globalState), _171_v2);
          _213_v32 = (_213_v32).update(new BigNumber((_214_v33).cardinality()), _172_v3);
          _170_v1 = _170_v1;
          let _215_v34;
          _215_v34 = _module.D4.create_DC12(new BigNumber((_182_v12).length), _169_v0, _182_v12, new BigNumber(350), new BigNumber((_174_v5).length));
          _187_v17 = (_187_v17).update(_dafny.Seq.IsProperPrefixOf(_182_v12, _dafny.Seq.UnicodeFromString("fvye")), (_215_v34).dtor_cf17);
          let _216_v35;
          let _nw31 = Array((new BigNumber(24)).toNumber()).fill(false);
          _216_v35 = _nw31;
          let _index20 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_216_v35).length));
          (_216_v35)[_index20] = _171_v2;
          let _index21 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_216_v35).length));
          (_216_v35)[_index21] = _171_v2;
          (_181_globalState).f2 = _171_v2;
        } else {
          (_181_globalState).f2 = _171_v2;
          _169_v0 = new BigNumber((_dafny.Seq.Concat(_182_v12, _dafny.Seq.UnicodeFromString("poavktqf"))).length);
          let _217_v36;
          let _init6 = ((_218_v6) => function (_219_i2) {
            return _218_v6;
          })(_175_v6);
          let _nw32 = Array((new BigNumber(3)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw32.length); _i0_6++) {
            _nw32[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _217_v36 = _nw32;
          let _index22 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_217_v36).length));
          (_217_v36)[_index22] = _dafny.Map.Empty.slice().updateUnsafe(_169_v0,_171_v2);
          let _220_v37;
          _220_v37 = _module.D8.create_DC21(_171_v2, _171_v2, _dafny.Seq.update(_dafny.Seq.update(_174_v5, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("wcd")).length), new BigNumber((_174_v5).length)), _171_v2), _module.__default.safeIndex(new BigNumber(-577), new BigNumber((_dafny.Seq.update(_174_v5, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("wcd")).length), new BigNumber((_174_v5).length)), _171_v2)).length)), _171_v2), _171_v2, _169_v0);
          let _index23 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_217_v36).length));
          (_217_v36)[_index23] = (_175_v6).update(_169_v0, (_220_v37).dtor_cf27);
          _174_v5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_174_v5, _module.__default.safeIndex(_169_v0, new BigNumber((_174_v5).length)), _171_v2), _174_v5), _174_v5);
          let _221_v38;
          let _nw33 = new _module.C4();
          _nw33.__ctor(_169_v0);
          _221_v38 = _nw33;
          let _222_v39;
          _222_v39 = _dafny.Seq.of(_221_v38, _221_v38);
          _module.__default.m0(_171_v2, _189_v18, _module.__default.safeDivisionInt(_169_v0, _module.__default.fm0(_169_v0, _169_v0, _169_v0, _171_v2, _181_globalState)), (_187_v17).update(_171_v2, new BigNumber((_222_v39).length)), _181_globalState);
        }
        let _223_v40;
        _223_v40 = _dafny.Map.Empty.slice().updateUnsafe(_169_v0,_module.D8.create_DC21(_171_v2, _171_v2, _174_v5, _171_v2, _169_v0));
        _223_v40 = _module.__default.fm33(_171_v2, _189_v18, _181_globalState);
        let _224_v41;
        let _init7 = ((_225_v18) => function (_226_i3) {
          return _225_v18;
        })(_189_v18);
        let _nw34 = Array((new BigNumber(25)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw34.length); _i0_7++) {
          _nw34[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _224_v41 = _nw34;
        let _index24 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_224_v41).length));
        (_224_v41)[_index24] = _189_v18;
        let _index25 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_224_v41).length));
        (_224_v41)[_index25] = _189_v18;
      } else {
        let _227_v42;
        _227_v42 = _dafny.Map.Empty.slice().updateUnsafe(_169_v0,_169_v0);
        _227_v42 = (_227_v42).update(_169_v0, _169_v0);
        _169_v0 = (_dafny.ZERO).minus(new BigNumber((_187_v17).length));
        let _228_v43;
        let _nw35 = new _module.C1();
        _nw35.__ctor();
        _228_v43 = _nw35;
        _228_v43 = _228_v43;
        (_181_globalState).f2 = _171_v2;
        (_181_globalState).f2 = _171_v2;
      }
      (_181_globalState).f2 = (_169_v0).isLessThan(_169_v0);
      let _229_v44;
      _229_v44 = _dafny.Map.Empty.slice().updateUnsafe(_182_v12,_169_v0);
      _229_v44 = (_229_v44).update(_182_v12, _module.__default.safeDivisionInt(_169_v0, _169_v0));
      _module.__default.m0(_171_v2, _189_v18, _169_v0, (_187_v17).Merge(_187_v17), _181_globalState);
      let _230_v45;
      _230_v45 = _dafny.MultiSet.fromElements(_171_v2);
      (_181_globalState).f2 = (_230_v45).IsDisjointFrom(_230_v45);
      if (_171_v2) {
        let _231_v46;
        _231_v46 = _dafny.MultiSet.fromElements(_189_v18, _189_v18);
        if ((_module.__default.fm0((((_231_v46).contains(_189_v18)) ? ((_231_v46).get(_189_v18)) : (_169_v0)), new BigNumber((_170_v1).length), _169_v0, _171_v2, _181_globalState)).isLessThanOrEqualTo(_169_v0)) {
          _module.__default.m0(_171_v2, _189_v18, _169_v0, _187_v17, _181_globalState);
          let _232_v47;
          let _nw36 = new _module.C3();
          _nw36.__ctor(new BigNumber(997));
          _232_v47 = _nw36;
          let _233_v48;
          _233_v48 = _dafny.Set.fromElements(_232_v47);
          let _234_v49;
          _234_v49 = _module.D8.create_DC20((_233_v48).Intersect(_233_v48));
          _234_v49 = (((_169_v0).isLessThan(_169_v0)) ? (_234_v49) : (_234_v49));
          let _235_v50;
          let _nw37 = Array((new BigNumber(25)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _171_v2;
          _nw37[(_dafny.ONE).toNumber()] = _171_v2;
          _nw37[(new BigNumber(2)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(3)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(4)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(5)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(6)).toNumber()] = false;
          _nw37[(new BigNumber(7)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(8)).toNumber()] = true;
          _nw37[(new BigNumber(9)).toNumber()] = true;
          _nw37[(new BigNumber(10)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(11)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(12)).toNumber()] = true;
          _nw37[(new BigNumber(13)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(14)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(15)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(16)).toNumber()] = (_232_v47).fm4(_177_v9, _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(793),_171_v2)), _dafny.MultiSet.FromArray(_dafny.Seq.of(_171_v2)), _181_globalState);
          _nw37[(new BigNumber(17)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(18)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(19)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(20)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(21)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(22)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(23)).toNumber()] = _171_v2;
          _nw37[(new BigNumber(24)).toNumber()] = _171_v2;
          _235_v50 = _nw37;
          let _236_v51;
          _236_v51 = _dafny.Map.Empty.slice().updateUnsafe(_170_v1,_235_v50);
          let _237_v52;
          _237_v52 = _dafny.Map.Empty.slice().updateUnsafe(_236_v51,(_module.D9.create_DC24(_171_v2, false, _171_v2)).dtor_cf34);
          _237_v52 = (_237_v52).update(_236_v51, (_180_v11).contains(_171_v2));
          (_181_globalState).f0 = true;
          (_181_globalState).f2 = _module.__default.fm1(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_238_v18) => function (_239_i4) {
            return _238_v18;
          })(_189_v18)), _module.__default.safeIndex(_169_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_240_v18) => function (_241_i4) {
            return _240_v18;
          })(_189_v18))).length)), _189_v18), _182_v12), _181_globalState);
        } else {
          let _rhs20 = _module.__default.fm0(_169_v0, _169_v0, (_dafny.ZERO).minus(_169_v0), !(_171_v2), _181_globalState);
          let _rhs21 = (((!(_171_v2)) ? (_231_v46) : (_231_v46))).IsSubsetOf(_dafny.MultiSet.fromElements(_189_v18));
          let _lhs19 = _181_globalState;
          _169_v0 = _rhs20;
          _lhs19.f2 = _rhs21;
          let _242_v53;
          _242_v53 = _dafny.Map.Empty.slice().updateUnsafe(_171_v2,_171_v2);
          let _243_v54;
          _243_v54 = _module.D2.create_DC7(new BigNumber((_182_v12).length), _dafny.Map.Empty.slice().updateUnsafe(_242_v53,_169_v0), _169_v0);
          let _244_v55;
          _244_v55 = _dafny.MultiSet.fromElements(_243_v54);
          let _245_v56;
          _245_v56 = _dafny.Map.Empty.slice().updateUnsafe((((_244_v55).contains(_243_v54)) ? ((_244_v55).get(_243_v54)) : (_169_v0)),_169_v0);
          let _246_v57;
          let _nw38 = new _module.C1();
          _nw38.__ctor();
          _246_v57 = _nw38;
          let _247_v58;
          _247_v58 = _dafny.Set.fromElements(_246_v57);
          _245_v56 = (_245_v56).update(_module.__default.safeModuloInt(_169_v0, _169_v0), new BigNumber(((_247_v58).Intersect(_247_v58)).length));
          _172_v3 = _172_v3;
          (_181_globalState).f5 = _169_v0;
          let _248_v59;
          let _nw39 = new _module.C3();
          _nw39.__ctor(_169_v0);
          _248_v59 = _nw39;
          let _249_v60;
          _249_v60 = _module.D2.create_DC5(_248_v59);
          let _250_v61;
          _250_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_174_v5),(_249_v60).dtor_cf5);
          let _251_v62;
          _251_v62 = _module.D9.create_DC24(_171_v2, _171_v2, _171_v2);
          let _252_v63;
          let _nw40 = Array((new BigNumber(15)).toNumber());
          _nw40[(_dafny.ZERO).toNumber()] = _248_v59;
          _nw40[(_dafny.ONE).toNumber()] = ((_171_v2) ? (_248_v59) : (_248_v59));
          _nw40[(new BigNumber(2)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(3)).toNumber()] = (((_250_v61).contains(_230_v45)) ? ((_250_v61).get(_230_v45)) : (_248_v59));
          _nw40[(new BigNumber(4)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(5)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(6)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(7)).toNumber()] = (((_251_v62).dtor_cf35) ? (_248_v59) : (_248_v59));
          _nw40[(new BigNumber(8)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(9)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(10)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(11)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(12)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(13)).toNumber()] = _248_v59;
          _nw40[(new BigNumber(14)).toNumber()] = _248_v59;
          _252_v63 = _nw40;
          let _index26 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_252_v63).length));
          (_252_v63)[_index26] = _248_v59;
          let _253_v64;
          _253_v64 = _module.D13.create_DC38(_187_v17);
          let _254_v65;
          _254_v65 = _dafny.MultiSet.fromElements(_253_v64, _253_v64);
          let _255_v66;
          let _nw41 = Array((new BigNumber(3)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = !(_171_v2);
          _nw41[(_dafny.ONE).toNumber()] = (_254_v65).IsProperSubsetOf(_254_v65);
          _nw41[(new BigNumber(2)).toNumber()] = _171_v2;
          _255_v66 = _nw41;
          let _index27 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_255_v66).length));
          (_255_v66)[_index27] = _171_v2;
          let _256_v67;
          let _nw42 = new _module.C3();
          _nw42.__ctor(_169_v0);
          _256_v67 = _nw42;
          let _index28 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_252_v63).length));
          let _index29 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_255_v66).length));
          let _rhs22 = _248_v59;
          let _rhs23 = _171_v2;
          let _rhs24 = _256_v67;
          let _lhs20 = _252_v63;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_252_v63).length));
          let _lhs22 = _255_v66;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_255_v66).length));
          _lhs20[_lhs21] = _rhs22;
          _lhs22[_lhs23] = _rhs23;
          _256_v67 = _rhs24;
        }
        let _257_v68;
        let _nw43 = Array((new BigNumber(5)).toNumber());
        _257_v68 = _nw43;
        let _258_v69;
        let _nw44 = new _module.C2();
        _nw44.__ctor(_171_v2);
        _258_v69 = _nw44;
        let _index30 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_257_v68).length));
        (_257_v68)[_index30] = _258_v69;
        let _index31 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_257_v68).length));
        let _rhs25 = ((_171_v2) ? (_182_v12) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(938)), ((_259_v2, _260_v0, _261_v18, _262_v45) => function (_263_i5) {
          return (_module.D10.create_DC27(_259_v2, _260_v0, _261_v18, _262_v45)).dtor_cf40;
        })(_171_v2, _169_v0, _189_v18, _230_v45))));
        let _rhs26 = _258_v69;
        let _lhs24 = _257_v68;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_257_v68).length));
        _182_v12 = _rhs25;
        _lhs24[_lhs25] = _rhs26;
        (_181_globalState).f0 = (((_169_v0).isLessThanOrEqualTo(_169_v0)) ? (_171_v2) : (_171_v2));
        (_181_globalState).f5 = (_dafny.ZERO).minus(_169_v0);
        (_181_globalState).f5 = new BigNumber(((_module.__default.fm34(_169_v0, _169_v0, !(_171_v2), _181_globalState)).update(_180_v11, new BigNumber(-582))).length);
      } else {
        _module.__default.m0(!(_171_v2) || (_171_v2), _189_v18, _169_v0, _187_v17, _181_globalState);
        _171_v2 = !_dafny.Seq.contains(_170_v1, _169_v0);
        let _264_v70;
        let _nw45 = Array((new BigNumber(28)).toNumber());
        _nw45[(_dafny.ZERO).toNumber()] = false;
        _nw45[(_dafny.ONE).toNumber()] = _171_v2;
        _nw45[(new BigNumber(2)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(3)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(4)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(5)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(6)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(7)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(8)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(9)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(10)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(11)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(12)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(13)).toNumber()] = !(_171_v2);
        _nw45[(new BigNumber(14)).toNumber()] = _module.__default.fm1(!(_171_v2), _181_globalState);
        _nw45[(new BigNumber(15)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(16)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(17)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(18)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(19)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(20)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(21)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(22)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(23)).toNumber()] = !(_171_v2);
        _nw45[(new BigNumber(24)).toNumber()] = _171_v2;
        _nw45[(new BigNumber(25)).toNumber()] = false;
        _nw45[(new BigNumber(26)).toNumber()] = true;
        _nw45[(new BigNumber(27)).toNumber()] = _171_v2;
        _264_v70 = _nw45;
        let _265_v71;
        let _nw46 = new _module.C0();
        _nw46.__ctor(true, _264_v70);
        _265_v71 = _nw46;
        let _266_v72;
        let _nw47 = new _module.C2();
        _nw47.__ctor((_265_v71).f10);
        _266_v72 = _nw47;
        let _267_v73;
        let _nw48 = Array((new BigNumber(6)).toNumber()).fill([]);
        _267_v73 = _nw48;
        let _268_v74;
        let _269_v75;
        let _270_v76;
        let _271_v77;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = (_266_v72).m1(_169_v0, _267_v73, new BigNumber(((_dafny.MultiSet.FromArray(_174_v5)).Union(_dafny.MultiSet.FromArray(_174_v5))).cardinality()), _169_v0, _181_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _268_v74 = _out0;
        _269_v75 = _out1;
        _270_v76 = _out2;
        _271_v77 = _out3;
      }
      let _272_v78;
      let _nw49 = Array((new BigNumber(10)).toNumber()).fill(false);
      _272_v78 = _nw49;
      let _273_v79;
      let _nw50 = new _module.C0();
      _nw50.__ctor(_171_v2, _272_v78);
      _273_v79 = _nw50;
      let _274_v80;
      _274_v80 = _dafny.Set.fromElements(_273_v79);
      let _275_v81;
      _275_v81 = _dafny.Map.Empty.slice().updateUnsafe(_171_v2,_274_v80);
      let _276_v82;
      _276_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_275_v81).contains((_273_v79).fm4(_177_v9, _module.D0.create_DC0(_175_v6), _230_v45, _181_globalState))) ? ((_275_v81).get((_273_v79).fm4(_177_v9, _module.D0.create_DC0(_175_v6), _230_v45, _181_globalState))) : (_274_v80))).length),_169_v0);
      let _277_v83;
      let _nw51 = new _module.C0();
      _nw51.__ctor((_273_v79).f10, (_273_v79).f11);
      _277_v83 = _nw51;
      let _278_v84;
      _278_v84 = _dafny.Map.Empty.slice().updateUnsafe(_277_v83,_169_v0);
      let _279_v85;
      _279_v85 = _dafny.Map.Empty.slice().updateUnsafe(_171_v2,_278_v84);
      _276_v82 = (_276_v82).update((((_273_v79).f10) ? (_169_v0) : (_module.__default.fm0(_169_v0, _169_v0, new BigNumber((_279_v85).length), (_273_v79).f10, _181_globalState))), _169_v0);
      let _280_v86;
      let _init8 = ((_281_v2, _282_v0) => function (_283_i6) {
        return _module.D13.create_DC38(_dafny.Map.Empty.slice().updateUnsafe(_281_v2,_282_v0));
      })(_171_v2, _169_v0);
      let _nw52 = Array((new BigNumber(10)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw52.length); _i0_8++) {
        _nw52[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _280_v86 = _nw52;
      let _284_v87;
      _284_v87 = _dafny.Set.fromElements(_280_v86);
      if ((_284_v87).IsDisjointFrom(_284_v87)) {
        _180_v11 = _module.__default.fm17(!((_273_v79).f10), (_dafny.ZERO).minus(_169_v0), (_273_v79).f10, _169_v0, _181_globalState);
        let _285_v88;
        let _nw53 = new _module.C4();
        _nw53.__ctor(_169_v0);
        _285_v88 = _nw53;
        let _286_v89;
        _286_v89 = _dafny.Seq.of(_285_v88);
        let _287_v90;
        _287_v90 = _dafny.Set.fromElements(new BigNumber((_286_v89).length), (_285_v88).f9, (_285_v88).f9);
        let _288_v91;
        _288_v91 = _dafny.Map.Empty.slice().updateUnsafe(_273_v79,new BigNumber((_287_v90).length));
        _288_v91 = (_288_v91).update(_273_v79, (_dafny.ZERO).minus(_169_v0));
        let _nw54 = Array((new BigNumber(13)).toNumber()).fill(false);
        _272_v78 = _nw54;
        let _289_v92;
        _289_v92 = _module.D3.create_DC8(_182_v12);
        (_181_globalState).f2 = _dafny.Seq.IsPrefixOf((_289_v92).dtor_cf10, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(956)), function (_290_i7) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }), _182_v12));
        let _291_v93;
        let _nw55 = new _module.C1();
        _nw55.__ctor();
        _291_v93 = _nw55;
        let _292_v94;
        _292_v94 = _dafny.Map.Empty.slice().updateUnsafe(_291_v93,_171_v2);
        _169_v0 = (new BigNumber((_292_v94).length)).multipliedBy((_285_v88).f9);
      } else {
        (_181_globalState).f2 = !(_169_v0).isEqualTo(new BigNumber((_dafny.Seq.update(_182_v12, _module.__default.safeIndex(_169_v0, new BigNumber((_182_v12).length)), _189_v18)).length));
        let _293_v95;
        let _nw56 = new _module.C4();
        _nw56.__ctor(_169_v0);
        _293_v95 = _nw56;
        _171_v2 = _171_v2;
        let _294_v96;
        let _nw57 = new _module.C1();
        _nw57.__ctor();
        _294_v96 = _nw57;
        (_181_globalState).f5 = (_169_v0).plus(_169_v0);
      }
      let _295_v97;
      let _nw58 = new _module.C1();
      _nw58.__ctor();
      _295_v97 = _nw58;
      let _296_v98;
      _296_v98 = _dafny.Seq.of(_295_v97, _295_v97);
      _295_v97 = (_296_v98)[_module.__default.safeIndex(_169_v0, new BigNumber((_296_v98).length))];
      process.stdout.write(_dafny.toString(_169_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_170_v1, _dafny.Seq.of(new BigNumber(292), new BigNumber(292)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_171_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_173_v4).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_174_v5, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_176_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(292),_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_177_v9).dtor_cf0).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false).updateUnsafe(new BigNumber(2),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v10).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false).updateUnsafe(new BigNumber(2),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v11).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_181_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_181_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_181_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_181_globalState.f3).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_181_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_181_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_181_globalState).f6).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false).updateUnsafe(new BigNumber(2),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_181_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_181_globalState.f8).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_182_v12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_183_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(7)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_189_v18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v20).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_193_v21).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D0.create_DC1(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v44).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("aiwqmuo"),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v45).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_273_v79).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_274_v80).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_275_v81).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v82).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(-1)).updateUnsafe(new BigNumber(-1),new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_278_v84).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_279_v85).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[_dafny.ZERO]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[_dafny.ONE]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(2)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(3)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(4)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(5)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(6)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(7)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(8)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_280_v86)[new BigNumber(9)]).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_284_v87).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_296_v98).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D0(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D1(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf3 === other.cf3;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf6) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC7(cf7, cf8, cf9) {
      let $dt = new D2(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC5(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf6 === other.cf6;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf5 === other.cf5;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf11) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC8(cf10) {
      let $dt = new D3(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC10(cf12) {
      let $dt = new D3(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + this.cf10.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf14, cf15, cf16, cf17, cf18) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC13(cf19, cf20) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC14() {
      let $dt = new D4(2);
      return $dt;
    }
    static create_DC11(cf13) {
      let $dt = new D4(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + this.cf16.toVerbatimString(true) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC14";
      } else if (this.$tag === 3) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf13 === other.cf13;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(_dafny.ZERO, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf21) {
      let $dt = new D5(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf22) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC17(cf23) {
      let $dt = new D7(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC19(cf24) {
      let $dt = new D7(2);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18";
      } else if (this.$tag === 1) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf26, cf27, cf28, cf29, cf30) {
      let $dt = new D8(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC20(cf25) {
      let $dt = new D8(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC22(cf31) {
      let $dt = new D8(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && this.cf27 === other.cf27 && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(false, false, _dafny.Seq.of(), false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf33, cf34, cf35) {
      let $dt = new D9(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC23(cf32) {
      let $dt = new D9(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC25(cf36) {
      let $dt = new D9(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33 && this.cf34 === other.cf34 && this.cf35 === other.cf35;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24(false, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf38, cf39, cf40, cf41) {
      let $dt = new D10(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC26(cf37) {
      let $dt = new D10(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC28(cf42) {
      let $dt = new D10(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39) && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC27(false, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.MultiSet.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30() {
      let $dt = new D11(0);
      return $dt;
    }
    static create_DC31(cf44, cf45) {
      let $dt = new D11(1);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC29(cf43) {
      let $dt = new D11(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC32(cf46) {
      let $dt = new D11(3);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get is_DC32() { return this.$tag === 3; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30";
      } else if (this.$tag === 1) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf43 === other.cf43;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC30();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC34(cf48, cf49, cf50) {
      let $dt = new D12(0);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC35() {
      let $dt = new D12(1);
      return $dt;
    }
    static create_DC36(cf51, cf52, cf53) {
      let $dt = new D12(2);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC33(cf47) {
      let $dt = new D12(3);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC37(cf54) {
      let $dt = new D12(4);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get is_DC37() { return this.$tag === 4; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC35";
      } else if (this.$tag === 2) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 4) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf48, other.cf48) && this.cf49 === other.cf49 && this.cf50 === other.cf50;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC34(_module.D0.Default(), false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC39() {
      let $dt = new D13(0);
      return $dt;
    }
    static create_DC40(cf56, cf57, cf58) {
      let $dt = new D13(1);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC38(cf55) {
      let $dt = new D13(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC39";
      } else if (this.$tag === 1) {
        return "D13.DC40" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf56 === other.cf56 && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC39();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf59) {
      let $dt = new D14(0);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf59 === other.cf59;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this.f2 = false;
      this.f3 = _dafny.MultiSet.Empty;
      this.f5 = _dafny.ZERO;
      this.f8 = _dafny.Set.Empty;
      this._f1 = _dafny.ZERO;
      this._f4 = false;
      this._f6 = _dafny.MultiSet.Empty;
      this._f7 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f10 = false;
      this._f11 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f10, f11) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f10;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      (globalState).f2 = (_this).f10;
      let _297_v0;
      _297_v0 = _dafny.Seq.UnicodeFromString("asdhidii");
      let _298_i0;
      _298_i0 = _dafny.ZERO;
      L1: {
        while (_dafny.Seq.IsPrefixOf(_297_v0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sqahgg"), _297_v0))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_298_i0)) {
              break L1;
            }
            _298_i0 = (_298_i0).plus(_dafny.ONE);
            let _299_v1;
            _299_v1 = _module.D0.create_DC1((_this).f10);
            let _300_v2;
            _300_v2 = _dafny.Map.Empty.slice().updateUnsafe(((_299_v1).dtor_cf1) === ((_this).f10),(_this).f10);
            _300_v2 = (_300_v2).update((_this).f10, (_this).f10);
            let _301_v3;
            let _nw59 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
            _301_v3 = _nw59;
            let _index32 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_301_v3).length));
            (_301_v3)[_index32] = p2;
            let _index33 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_301_v3).length));
            (_301_v3)[_index33] = p3;
            let _302_v4;
            _302_v4 = _dafny.Seq.of((_this).f10);
            let _rhs27 = (_this).f10;
            let _rhs28 = _dafny.Seq.Concat(_302_v4, _302_v4);
            let _rhs29 = (_this).f10;
            let _rhs30 = (_this).f10;
            let _lhs26 = globalState;
            let _lhs27 = globalState;
            r0 = _rhs27;
            _302_v4 = _rhs28;
            _lhs26.f2 = _rhs29;
            _lhs27.f2 = _rhs30;
            (globalState).f5 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_301_v3)[_module.__default.safeIndex(new BigNumber(128), new BigNumber((_301_v3).length))], new BigNumber((_297_v0).length)), p3);
          }
        }
      }
      let _303_v5;
      let _nw60 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
      _303_v5 = _nw60;
      let _304_v6;
      _304_v6 = _module.D1.create_DC3(_303_v5);
      let _305_v7;
      _305_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_303_v5);
      let _306_v8;
      _306_v8 = _dafny.Seq.of(_303_v5, _303_v5);
      let _307_v9;
      let _nw61 = Array((new BigNumber(21)).toNumber());
      _nw61[(_dafny.ZERO).toNumber()] = _303_v5;
      _nw61[(_dafny.ONE).toNumber()] = _303_v5;
      _nw61[(new BigNumber(2)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(3)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(4)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(5)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(6)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(7)).toNumber()] = (_304_v6).dtor_cf3;
      _nw61[(new BigNumber(8)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(9)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(10)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(11)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(12)).toNumber()] = (((_305_v7).contains((_this).f10)) ? ((_305_v7).get((_this).f10)) : ((_306_v8)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_306_v8).length))]));
      _nw61[(new BigNumber(13)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(14)).toNumber()] = (((_this).f10) ? (_303_v5) : (_303_v5));
      _nw61[(new BigNumber(15)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(16)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(17)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(18)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(19)).toNumber()] = _303_v5;
      _nw61[(new BigNumber(20)).toNumber()] = _303_v5;
      _307_v9 = _nw61;
      let _index34 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_307_v9).length));
      (_307_v9)[_index34] = _303_v5;
      let _index35 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_307_v9).length));
      let _rhs31 = _303_v5;
      let _rhs32 = _module.__default.fm0(new BigNumber((_297_v0).length), p2, p0, ((_this).f10) === (!((_this).f10)), globalState);
      let _rhs33 = p0;
      let _lhs28 = _307_v9;
      let _lhs29 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_307_v9).length));
      let _lhs30 = globalState;
      _lhs28[_lhs29] = _rhs31;
      _lhs30.f5 = _rhs32;
      r1 = _rhs33;
      let _308_v10;
      _308_v10 = _dafny.Set.fromElements((_this).f10);
      let _309_v11;
      _309_v11 = _dafny.Set.fromElements(new BigNumber((_308_v10).length));
      _309_v11 = _module.__default.fm6(globalState);
      if (!(true)) {
        (globalState).f0 = (((_this).f10) ? ((_this).f10) : (!(false)));
        let _310_v12;
        _310_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,!((_this).f10));
        let _311_v13;
        _311_v13 = _dafny.Seq.of(((((_310_v12).contains((_this).f10)) ? ((_310_v12).get((_this).f10)) : ((_this).f10))) === ((_this).f10));
        _311_v13 = _dafny.Seq.Concat(_dafny.Seq.of((_this).f10), _dafny.Seq.Concat(_311_v13, _311_v13));
        let _312_v14;
        let _init9 = ((_313_p3) => function (_314_i1) {
          return _module.__default.safeModuloInt(_314_i1, _313_p3);
        })(p3);
        let _nw62 = Array((new BigNumber(7)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw62.length); _i0_9++) {
          _nw62[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _312_v14 = _nw62;
        _312_v14 = _312_v14;
        r1 = p0;
        (globalState).f0 = (_this).f10;
      } else {
        let _315_v15;
        _315_v15 = _dafny.MultiSet.fromElements((_this).f10);
        (globalState).f2 = (_315_v15).IsProperSubsetOf((_module.__default.fm7(globalState)).Union((_315_v15).update((_this).f10, _module.__default.abs(p3))));
        r0 = true;
        let _316_v16;
        let _init10 = ((_317_p0) => function (_318_i2) {
          return (_318_i2).minus(_317_p0);
        })(p0);
        let _nw63 = Array((new BigNumber(13)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw63.length); _i0_10++) {
          _nw63[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _316_v16 = _nw63;
        _316_v16 = _316_v16;
        let _319_v17;
        _319_v17 = _dafny.Seq.of(false);
        let _index36 = _module.__default.safeIndex(new BigNumber(773), new BigNumber(((_this).f11).length));
        ((_this).f11)[_index36] = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_319_v17, _module.__default.safeIndex(p2, new BigNumber((_319_v17).length)), (_319_v17)[_module.__default.safeIndex(new BigNumber((_319_v17).length), new BigNumber((_319_v17).length))]), _319_v17);
        let _index37 = _module.__default.safeIndex(new BigNumber(773), new BigNumber(((_this).f11).length));
        ((_this).f11)[_index37] = (p2).isLessThanOrEqualTo(p3);
        if ((_this).f10) {
          (globalState).f5 = _module.__default.fm0(p3, p3, p2, ((_this).f11)[_module.__default.safeIndex(new BigNumber(773), new BigNumber(((_this).f11).length))], globalState);
          (globalState).f0 = _module.__default.fm1(false, globalState);
          let _320_v18;
          _320_v18 = _dafny.Map.Empty.slice().updateUnsafe(p3,((_this).f11)[_module.__default.safeIndex(new BigNumber(773), new BigNumber(((_this).f11).length))]);
          let _321_v19;
          _321_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_320_v18).length),_module.__default.safeDivisionInt(p2, new BigNumber(682)));
          _321_v19 = (_321_v19).update(p0, (((_321_v19).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("weoxwqp")).length)))) ? ((_321_v19).get((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("weoxwqp")).length)))) : (p2)));
          (globalState).f2 = (_this).f10;
          let _322_v20;
          _322_v20 = _dafny.Map.Empty.slice().updateUnsafe(_308_v10,(p3).isEqualTo(p3));
          _322_v20 = _322_v20;
        } else {
          let _323_v21;
          _323_v21 = new _dafny.CodePoint('h'.codePointAt(0));
          _323_v21 = _323_v21;
          let _324_v22;
          _324_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
          let _325_v23;
          _325_v23 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f11)[_module.__default.safeIndex(new BigNumber(773), new BigNumber(((_this).f11).length))],(_this).f10);
          let _326_v24;
          _326_v24 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f11)[_module.__default.safeIndex(new BigNumber(773), new BigNumber(((_this).f11).length))],((((_324_v22).contains(_module.__default.fm0(p2, new BigNumber((_325_v23).length), p2, (_this).f10, globalState))) ? ((_324_v22).get(_module.__default.fm0(p2, new BigNumber((_325_v23).length), p2, (_this).f10, globalState))) : (p0))).multipliedBy((_dafny.ZERO).minus(p0)));
          _326_v24 = (_326_v24).update((_this).f10, (p0).minus(p0));
          let _327_v25;
          _327_v25 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f10);
          let _index38 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_316_v16).length));
          (_316_v16)[_index38] = new BigNumber(-985);
          let _328_v26;
          _328_v26 = _module.D0.create_DC0(_327_v25);
          let _329_v27;
          _329_v27 = _dafny.Set.fromElements(_323_v21);
          let _pat_let_tv1 = _307_v9;
          let _pat_let_tv2 = _307_v9;
          let _index39 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_316_v16).length));
          let _rhs34 = (_328_v26).dtor_cf0;
          let _rhs35 = new BigNumber(699);
          let _rhs36 = function (_pat_let2_0) {
            return function (_330_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_331_dt__update_hcf3_h0) {
                  return _module.D1.create_DC3(_331_dt__update_hcf3_h0);
                }(_pat_let3_0);
              }((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_pat_let_tv1).length))]);
            }(_pat_let2_0);
          }(_304_v6);
          let _rhs37 = (((_329_v27).IsProperSubsetOf(_329_v27)) ? ((_this).f10) : (((_this).f11)[_module.__default.safeIndex(new BigNumber(773), new BigNumber(((_this).f11).length))]));
          let _lhs31 = _316_v16;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_316_v16).length));
          let _lhs33 = globalState;
          _327_v25 = _rhs34;
          _lhs31[_lhs32] = _rhs35;
          _304_v6 = _rhs36;
          _lhs33.f0 = _rhs37;
          let _332_v28;
          _332_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(p3),p2);
          let _333_v29;
          _333_v29 = _module.D1.create_DC4((_316_v16)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_316_v16).length))]);
          let _334_v30;
          _334_v30 = _dafny.MultiSet.fromElements(new BigNumber(675));
          _332_v28 = (_332_v28).update(_333_v29, _module.__default.fm0(new BigNumber(492), p2, new BigNumber((_334_v30).cardinality()), false, globalState));
          (globalState).f2 = true;
        }
      }
      let _335_v31;
      _335_v31 = _dafny.Seq.of(p0);
      let _336_v32;
      _336_v32 = _dafny.Map.Empty.slice().updateUnsafe((_335_v31)[_module.__default.safeIndex(p0, new BigNumber((_335_v31).length))],(_dafny.ZERO).minus(p0));
      r1 = (new BigNumber((_336_v32).length)).minus(((_dafny.ZERO).minus(p3)).minus(p3));
      r0 = (_this).f10;
      let _337_v33;
      _337_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
      let _338_v34;
      _338_v34 = _module.D0.create_DC0(_337_v33);
      let _339_v35;
      _339_v35 = _dafny.MultiSet.fromElements((_this).f10, (_this).f10);
      let _340_v36;
      _340_v36 = _dafny.MultiSet.fromElements((_this).f10, (_this).fm4(_338_v34, _338_v34, _339_v35, globalState), (_this).f10);
      let _341_v37;
      _341_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(_338_v34, _338_v34, (_339_v35).update((_this).f10, _module.__default.abs(p0)), globalState),p3);
      r1 = (((_341_v37).contains((_this).f10)) ? ((_341_v37).get((_this).f10)) : (p3));
      r2 = ((_this).f10) && ((((_337_v33).contains(p3)) ? ((_337_v33).get(p3)) : ((_this).fm4(_338_v34, _338_v34, _340_v36, globalState))));
      let _342_v38;
      _342_v38 = new _dafny.CodePoint('w'.codePointAt(0));
      let _343_v39;
      _343_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_342_v38);
      r3 = _dafny.Seq.IsProperPrefixOf(_module.__default.fm8((_this).f10, _343_v39, (_this).f10, globalState), _297_v0);
      return [r0, r1, r2, r3];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      let _source5 = _module.D1.create_DC4(new BigNumber(-738));
      if (_source5.is_DC4) {
        let _344___mcc_h0 = (_source5).cf4;
        let _345_cf4 = _344___mcc_h0;
        return true;
      } else {
        let _346___mcc_h1 = (_source5).cf3;
        let _347_cf3 = _346___mcc_h1;
        return true;
      }
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _348_v0;
      _348_v0 = true;
      let _349_v1;
      let _nw64 = Array((new BigNumber(9)).toNumber()).fill(false);
      _349_v1 = _nw64;
      let _350_v2;
      let _nw65 = new _module.C0();
      _nw65.__ctor(_348_v0, _349_v1);
      _350_v2 = _nw65;
      let _index40 = _module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length));
      ((_350_v2).f11)[_index40] = true;
      let _351_v3;
      _351_v3 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_350_v2).f10);
      let _352_v4;
      _352_v4 = _module.D0.create_DC0(_351_v3);
      let _353_v5;
      _353_v5 = _module.D0.create_DC0((_352_v4).dtor_cf0);
      let _354_v6;
      _354_v6 = _dafny.MultiSet.fromElements(_module.__default.fm1(_348_v0, globalState));
      let _355_v7;
      _355_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,_354_v6);
      let _index41 = _module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length));
      ((_350_v2).f11)[_index41] = (_350_v2).fm4(_352_v4, _353_v5, (((_355_v7).contains(p3)) ? ((_355_v7).get(p3)) : (_354_v6)), globalState);
      let _356_v8;
      _356_v8 = _dafny.Seq.of(p0);
      let _357_v9;
      _357_v9 = _356_v8;
      (globalState).f5 = _module.__default.fm0(p3, p2, new BigNumber(474), _dafny.Seq.IsPrefixOf((_357_v9), _356_v8), globalState);
      let _hi1 = p0;
      for (let _358_i0 = p2; _358_i0.isLessThan(_hi1); _358_i0 = _358_i0.plus(_dafny.ONE)) {
        let _359_v10;
        _359_v10 = _module.D7.create_DC17(_dafny.Seq.of((_350_v2).f10));
        let _360_v11;
        _360_v11 = _dafny.Seq.of(((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))], (_350_v2).f10, ((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))], (_this).fm4(_module.D0.create_DC0(_module.__default.fm22((_350_v2).f10, globalState)), _352_v4, _354_v6, globalState), true);
        let _361_v12;
        _361_v12 = _dafny.Set.fromElements(_360_v11);
        if ((_361_v12).contains((_359_v10).dtor_cf23)) {
          (globalState).f0 = (_module.D2.create_DC6((_350_v2).f10)).dtor_cf6;
          r2 = (p0).isLessThanOrEqualTo(p3);
          let _362_v13;
          _362_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_348_v0,_348_v0),p0);
          let _363_v14;
          _363_v14 = _module.D2.create_DC7(new BigNumber((_360_v11).length), _362_v13, _358_i0);
          (globalState).f5 = (_363_v14).dtor_cf9;
          let _364_v15;
          _364_v15 = _dafny.MultiSet.fromElements(new BigNumber(493));
          _364_v15 = _module.__default.fm23((_364_v15).update(_module.__default.fm0(p0, p3, p2, (_350_v2).f10, globalState), _module.__default.abs(p3)), p3, globalState);
          let _rhs38 = p2;
          let _rhs39 = _348_v0;
          let _rhs40 = (_350_v2).f11;
          let _rhs41 = _349_v1;
          let _rhs42 = new BigNumber(((_351_v3).Merge(_351_v3)).length);
          let _lhs34 = globalState;
          r1 = _rhs38;
          _lhs34.f0 = _rhs39;
          _349_v1 = _rhs40;
          _349_v1 = _rhs41;
          r1 = _rhs42;
        } else {
          let _365_v16;
          _365_v16 = _dafny.Map.Empty.slice().updateUnsafe(p3,_349_v1);
          _365_v16 = ((_365_v16).Merge(_365_v16)).Merge(_365_v16);
          r2 = _348_v0;
          let _366_v17;
          let _nw66 = new _module.C0();
          _nw66.__ctor((_350_v2).f10, (_350_v2).f11);
          _366_v17 = _nw66;
          let _367_v18;
          _367_v18 = _dafny.Set.fromElements(_366_v17);
          let _368_v19;
          _368_v19 = _module.D8.create_DC20(_367_v18);
          let _rhs43 = ((!((_350_v2).f10)) ? ((_350_v2).f11) : ((_350_v2).f11));
          let _rhs44 = ((!((_350_v2).f10)) ? ((_367_v18).IsSubsetOf((_368_v19).dtor_cf25)) : ((_350_v2).f10));
          let _rhs45 = (_350_v2).f10;
          let _lhs35 = globalState;
          let _lhs36 = globalState;
          _349_v1 = _rhs43;
          _lhs35.f2 = _rhs44;
          _lhs36.f2 = _rhs45;
          let _369_v20;
          let _nw67 = new _module.C0();
          _nw67.__ctor(((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))], (_350_v2).f11);
          _369_v20 = _nw67;
          r0 = ((_350_v2).f10) && (true);
        }
        let _370_v21;
        let _nw68 = new _module.C0();
        _nw68.__ctor(((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))], (_350_v2).f11);
        _370_v21 = _nw68;
        (globalState).f5 = new BigNumber((_dafny.MultiSet.fromElements((_350_v2).f10)).cardinality());
        let _371_v22;
        let _nw69 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _371_v22 = _nw69;
        let _372_v23;
        _372_v23 = _371_v22;
        let _373_v24;
        _373_v24 = _dafny.Map.Empty.slice().updateUnsafe((((_350_v2).f10) ? (_372_v23) : (_372_v23)),(_module.D4.create_DC13(p3, _349_v1)).dtor_cf19);
        _373_v24 = (_373_v24).update(_372_v23, _module.__default.safeDivisionInt(p0, p3));
      }
      let _374_v25;
      _374_v25 = _module.D2.create_DC6(false);
      let _375_i1;
      _375_i1 = _dafny.ZERO;
      L2: {
        while ((_374_v25).dtor_cf6) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_375_i1)) {
              break L2;
            }
            _375_i1 = (_375_i1).plus(_dafny.ONE);
            let _376_v26;
            let _nw70 = new _module.C0();
            _nw70.__ctor(_348_v0, _349_v1);
            _376_v26 = _nw70;
            (globalState).f2 = !(_module.__default.fm0(p0, p0, p2, (_350_v2).f10, globalState)).isEqualTo(new BigNumber((_351_v3).length));
            _351_v3 = (_351_v3).update(p3, ((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))]);
            let _377_v27;
            _377_v27 = new _dafny.CodePoint('m'.codePointAt(0));
            _377_v27 = _377_v27;
          }
        }
      }
      let _378_v28;
      _378_v28 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_356_v8));
      let _pat_let_tv3 = _348_v0;
      let _pat_let_tv4 = p0;
      let _pat_let_tv5 = _350_v2;
      let _pat_let_tv6 = _350_v2;
      let _pat_let_tv7 = _350_v2;
      let _pat_let_tv8 = _350_v2;
      let _pat_let_tv9 = p0;
      let _rhs46 = !(!(_348_v0) || (_348_v0));
      let _rhs47 = _module.__default.safeDivisionInt(p3, new BigNumber(652));
      let _rhs48 = ((_354_v6).Union(_354_v6)).IsDisjointFrom((_354_v6).update(_348_v0, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_378_v28)).cardinality())))));
      let _rhs49 = function (_source6) {
        if (_source6.is_DC1) {
          let _379___mcc_h0 = (_source6).cf1;
          let _380_cf1 = _379___mcc_h0;
          return new BigNumber(527);
        } else if (_source6.is_DC0) {
          let _381___mcc_h1 = (_source6).cf0;
          let _382_cf0 = _381___mcc_h1;
          return (new BigNumber((_dafny.Set.fromElements(_pat_let_tv3)).length)).multipliedBy(_pat_let_tv4);
        } else {
          let _383___mcc_h2 = (_source6).cf2;
          let _384_cf2 = _383___mcc_h2;
          return _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_pat_let_tv6).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_pat_let_tv5).f11).length))],((_pat_let_tv8).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_pat_let_tv7).f11).length))])).length), _pat_let_tv9);
        }
      }(_352_v4);
      let _rhs50 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p0, p2), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p3, p2)));
      let _lhs37 = globalState;
      r3 = _rhs46;
      _lhs37.f5 = _rhs47;
      _348_v0 = _rhs48;
      r1 = _rhs49;
      r1 = _rhs50;
      let _385_v29;
      _385_v29 = _dafny.Map.Empty.slice().updateUnsafe(((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))],false);
      let _386_v30;
      _386_v30 = _dafny.Map.Empty.slice().updateUnsafe(_385_v29,p0);
      let _387_v31;
      _387_v31 = _dafny.Seq.of(((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))]);
      let _388_v32;
      _388_v32 = _module.D2.create_DC7(p3, _386_v30, new BigNumber((_387_v31).length));
      r0 = ((_388_v32).dtor_cf9).isLessThan(p3);
      r1 = p2;
      r2 = !((_dafny.Map.Empty.slice().updateUnsafe(_348_v0,new BigNumber((_387_v31).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).fm4(_module.D0.create_DC0(_351_v3), _353_v5, _354_v6, globalState),_module.__default.fm0(p2, p0, p2, ((_350_v2).f11)[_module.__default.safeIndex(new BigNumber(154), new BigNumber(((_350_v2).f11).length))], globalState)))).contains(true);
      r3 = _348_v0;
      return [r0, r1, r2, r3];
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      (globalState).f0 = p2;
      let _389_v0;
      _389_v0 = _dafny.Seq.UnicodeFromString("igpdhpy");
      let _390_v1;
      _390_v1 = _dafny.Set.fromElements(_389_v0);
      _390_v1 = _390_v1;
      if (p2) {
        let _391_v2;
        _391_v2 = _module.D2.create_DC6(p2);
        let _392_v3;
        _392_v3 = _dafny.Seq.of((_391_v2).dtor_cf6, p2, p0, p2, !(p0));
        let _393_v4;
        _393_v4 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(new BigNumber((_392_v3).length), p1));
        let _394_v5;
        _394_v5 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _395_v6;
        let _nw71 = Array((new BigNumber(11)).toNumber());
        _nw71[(_dafny.ZERO).toNumber()] = p0;
        _nw71[(_dafny.ONE).toNumber()] = ((((_394_v5).contains(p2)) ? ((_394_v5).get(p2)) : (p0))) && (p2);
        _nw71[(new BigNumber(2)).toNumber()] = ((p2) ? (p0) : (p2));
        _nw71[(new BigNumber(3)).toNumber()] = true;
        _nw71[(new BigNumber(4)).toNumber()] = p2;
        _nw71[(new BigNumber(5)).toNumber()] = p2;
        _nw71[(new BigNumber(6)).toNumber()] = p0;
        _nw71[(new BigNumber(7)).toNumber()] = p0;
        _nw71[(new BigNumber(8)).toNumber()] = p0;
        _nw71[(new BigNumber(9)).toNumber()] = p2;
        _nw71[(new BigNumber(10)).toNumber()] = p2;
        _395_v6 = _nw71;
        let _396_v7;
        _396_v7 = _dafny.Map.Empty.slice().updateUnsafe((_392_v3)[_module.__default.safeIndex(new BigNumber((_393_v4).cardinality()), new BigNumber((_392_v3).length))],_392_v3);
        let _rhs51 = _393_v4;
        let _rhs52 = new BigNumber((((p0) ? (_389_v0) : (_dafny.Seq.UnicodeFromString("inuynen")))).length);
        let _rhs53 = !((!(p0)) && (p2));
        let _rhs54 = (new BigNumber((_396_v7).length)).isLessThanOrEqualTo((p1).multipliedBy(p1));
        let _rhs55 = _395_v6;
        let _lhs38 = globalState;
        let _lhs39 = globalState;
        let _lhs40 = globalState;
        _393_v4 = _rhs51;
        _lhs38.f5 = _rhs52;
        _lhs39.f0 = _rhs53;
        _lhs40.f0 = _rhs54;
        _395_v6 = _rhs55;
        let _397_v8;
        let _init11 = ((_398_p1) => function (_399_i0) {
          return (_399_i0).multipliedBy(_398_p1);
        })(p1);
        let _nw72 = Array((new BigNumber(10)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw72.length); _i0_11++) {
          _nw72[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _397_v8 = _nw72;
        let _index42 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length));
        (_397_v8)[_index42] = p1;
        let _400_v10;
        let _init12 = ((_401_p0, _402_p1) => function (_403_i1) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_401_p0,_dafny.Map.Empty.slice().updateUnsafe(_401_p0,new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(_402_p1), function () {
            let _coll21 = new _dafny.Set();
            for (const _compr_21 of (_dafny.Seq.of(_402_p1)).Elements) {
              let _404_v9 = _compr_21;
              if (_dafny.Seq.contains(_dafny.Seq.of(_402_p1), _404_v9)) {
                _coll21.add(_module.__default.safeModuloInt(_404_v9, new BigNumber(718)));
              }
            }
            return _coll21;
          }())).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_401_p0,_dafny.Map.Empty.slice().updateUnsafe(_401_p0,_402_p1)));
        })(p0, p1);
        let _nw73 = Array((new BigNumber(6)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw73.length); _i0_12++) {
          _nw73[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _400_v10 = _nw73;
        let _405_v11;
        _405_v11 = _dafny.Map.Empty.slice().updateUnsafe(_395_v6,p0);
        let _406_v12;
        _406_v12 = _module.D9.create_DC23(_405_v11);
        let _index43 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_400_v10).length));
        (_400_v10)[_index43] = _module.__default.fm24(new BigNumber(((_406_v12).dtor_cf32).length), globalState);
        let _407_v13;
        _407_v13 = _dafny.Set.fromElements(_module.__default.fm25((((_394_v5).contains(p2)) ? ((_394_v5).get(p2)) : (p2)), new _dafny.CodePoint('n'.codePointAt(0)), p0, globalState), _392_v3, _392_v3);
        let _408_v14;
        _408_v14 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
        let _409_v15;
        _409_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,_408_v14);
        let _index44 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length));
        let _index45 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_400_v10).length));
        let _rhs56 = new BigNumber((_407_v13).length);
        let _rhs57 = _392_v3;
        let _rhs58 = (_409_v15).Merge(_409_v15);
        let _lhs41 = _397_v8;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length));
        let _lhs43 = _400_v10;
        let _lhs44 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_400_v10).length));
        _lhs41[_lhs42] = _rhs56;
        _392_v3 = _rhs57;
        _lhs43[_lhs44] = _rhs58;
        let _410_v16;
        _410_v16 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bodi")).length), p1, p1, new BigNumber(-446), (_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))]);
        let _411_v17;
        _411_v17 = _dafny.Set.fromElements(!(_module.__default.fm1(p2, globalState)));
        if (!(((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))]).minus((_410_v16)[_module.__default.safeIndex(p1, new BigNumber((_410_v16).length))])).isEqualTo((p1).minus((_dafny.ZERO).minus(new BigNumber((_411_v17).length))))) {
          let _412_v18;
          _412_v18 = _module.D4.create_DC13((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], _395_v6);
          _412_v18 = _module.D4.create_DC13((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], _395_v6);
          (globalState).f2 = ((_dafny.ZERO).minus(p1)).isLessThan(((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))]).multipliedBy(new BigNumber((_394_v5).length)));
          let _413_v19;
          _413_v19 = _dafny.Map.Empty.slice().updateUnsafe((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))],_395_v6);
          _413_v19 = (_413_v19).update(p1, _395_v6);
          let _414_v20;
          _414_v20 = _dafny.Map.Empty.slice().updateUnsafe((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))],!(!((p2) && (p0))));
          let _pat_let_tv10 = p2;
          let _rhs59 = (((_414_v20).contains(_module.__default.fm0(new BigNumber((_411_v17).length), (_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], new BigNumber(-119), p2, globalState))) ? ((_414_v20).get(_module.__default.fm0(new BigNumber((_411_v17).length), (_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], new BigNumber(-119), p2, globalState))) : (p0));
          let _rhs60 = new BigNumber(437);
          let _rhs61 = (function (_pat_let4_0) {
            return function (_415_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_416_dt__update_hcf6_h0) {
                  return _module.D2.create_DC6(_416_dt__update_hcf6_h0);
                }(_pat_let5_0);
              }(_pat_let_tv10);
            }(_pat_let4_0);
          }(_391_v2)).dtor_cf6;
          let _rhs62 = _dafny.Seq.Concat(_389_v0, _dafny.Seq.Concat(_389_v0, _389_v0));
          let _lhs45 = globalState;
          let _lhs46 = globalState;
          let _lhs47 = globalState;
          _lhs45.f2 = _rhs59;
          _lhs46.f5 = _rhs60;
          _lhs47.f0 = _rhs61;
          _389_v0 = _rhs62;
          let _417_v21;
          let _nw74 = new _module.C0();
          _nw74.__ctor(p0, _395_v6);
          _417_v21 = _nw74;
        } else {
          let _index46 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_395_v6).length));
          (_395_v6)[_index46] = p2;
          let _418_v22;
          _418_v22 = _397_v8;
          let _index47 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_395_v6).length));
          let _index48 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length));
          let _rhs63 = _395_v6;
          let _rhs64 = (_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))];
          let _rhs65 = p0;
          let _rhs66 = p1;
          let _rhs67 = (_418_v22);
          let _lhs48 = globalState;
          let _lhs49 = _395_v6;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_395_v6).length));
          let _lhs51 = _397_v8;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length));
          _395_v6 = _rhs63;
          _lhs48.f5 = _rhs64;
          _lhs49[_lhs50] = _rhs65;
          _lhs51[_lhs52] = _rhs66;
          _397_v8 = _rhs67;
          let _419_v23;
          _419_v23 = _dafny.MultiSet.fromElements(((p0) ? ((_391_v2).dtor_cf6) : ((_392_v3)[_module.__default.safeIndex((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], new BigNumber((_392_v3).length))])));
          let _420_v24;
          _420_v24 = _module.D0.create_DC0(_module.__default.fm22(p2, globalState));
          let _rhs68 = (_dafny.ZERO).minus(new BigNumber((_419_v23).cardinality()));
          let _rhs69 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_module.__default.fm26((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], new BigNumber(901), false, (_this).fm4(_420_v24, _420_v24, _419_v23, globalState), globalState)).dtor_cf9), p1);
          let _lhs53 = globalState;
          let _lhs54 = globalState;
          _lhs53.f5 = _rhs68;
          _lhs54.f5 = _rhs69;
          let _421_v26;
          _421_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_389_v0).length),_393_v4);
          let _422_v27;
          _422_v27 = _dafny.Map.Empty.slice().updateUnsafe(((((_393_v4).contains(p1)) ? ((_393_v4).get(p1)) : ((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))]))).multipliedBy(new BigNumber((function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of _dafny.IntegerRange(new BigNumber(850), new BigNumber(991))) {
              let _423_v25 = _compr_22;
              if (((new BigNumber(850)).isLessThanOrEqualTo(_423_v25)) && ((_423_v25).isLessThan(new BigNumber(991)))) {
                _coll22.push([_module.__default.safeModuloInt(_423_v25, new BigNumber(416)),new BigNumber(132)]);
              }
            }
            return _coll22;
          }()).length)),((((_421_v26).contains(new BigNumber((_dafny.MultiSet.fromElements((_395_v6)[_module.__default.safeIndex(new BigNumber(481), new BigNumber((_395_v6).length))], p0)).cardinality()))) ? ((_421_v26).get(new BigNumber((_dafny.MultiSet.fromElements((_395_v6)[_module.__default.safeIndex(new BigNumber(481), new BigNumber((_395_v6).length))], p0)).cardinality()))) : (_393_v4))).Intersect(_393_v4));
          _422_v27 = (_422_v27).update((p1).minus(p1), _393_v4);
          let _index49 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_395_v6).length));
          (_395_v6)[_index49] = (((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))]).plus(p1)).isLessThan(_module.__default.safeModuloInt(p1, p1));
          let _424_v28;
          let _init13 = ((_425_p2, _426_p1, _427_v14) => function (_428_i2) {
            return ((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_425_p2,_426_p1)) : (_427_v14));
          })(p2, p1, _408_v14);
          let _nw75 = Array((_dafny.ONE).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw75.length); _i0_13++) {
            _nw75[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _424_v28 = _nw75;
          let _429_v29;
          let _nw76 = Array((new BigNumber(8)).toNumber()).fill(false);
          _429_v29 = _nw76;
          let _index50 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length));
          let _rhs70 = _424_v28;
          let _rhs71 = (_module.__default.fm0((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], p1, new BigNumber((_389_v0).length), !(p0), globalState)).plus(new BigNumber((_389_v0).length));
          let _rhs72 = (_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))];
          let _rhs73 = (p1).plus(new BigNumber((_dafny.Seq.of(_410_v16, _410_v16)).length));
          let _rhs74 = _429_v29;
          let _lhs55 = _397_v8;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length));
          let _lhs57 = globalState;
          _424_v28 = _rhs70;
          _lhs55[_lhs56] = _rhs71;
          r0 = _rhs72;
          _lhs57.f5 = _rhs73;
          _429_v29 = _rhs74;
        }
        let _430_v30;
        _430_v30 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,p2),(_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))]);
        let _431_v31;
        _431_v31 = _module.D2.create_DC7(p1, _430_v30, (_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))]);
        let _432_v32;
        let _nw77 = Array((new BigNumber(7)).toNumber());
        _nw77[(_dafny.ZERO).toNumber()] = _431_v31;
        _nw77[(_dafny.ONE).toNumber()] = _431_v31;
        _nw77[(new BigNumber(2)).toNumber()] = _431_v31;
        _nw77[(new BigNumber(3)).toNumber()] = _module.__default.fm26((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], (_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], p2, !(true), globalState);
        _nw77[(new BigNumber(4)).toNumber()] = _module.D2.create_DC7(p1, _430_v30, new BigNumber(587));
        _nw77[(new BigNumber(5)).toNumber()] = _module.D2.create_DC7((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], _430_v30, new BigNumber((_389_v0).length));
        _nw77[(new BigNumber(6)).toNumber()] = _module.__default.fm26(p1, p1, !(!(p0)), p2, globalState);
        _432_v32 = _nw77;
        let _index51 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_432_v32).length));
        (_432_v32)[_index51] = _module.D2.create_DC7((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], _430_v30, p1);
        let _index52 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_432_v32).length));
        (_432_v32)[_index52] = _431_v31;
        let _433_v33;
        _433_v33 = _module.D4.create_DC14();
        let _434_v34;
        _434_v34 = _dafny.Seq.of(_433_v33, _module.D4.create_DC14(), _433_v33);
        let _435_v35;
        _435_v35 = new _dafny.CodePoint('s'.codePointAt(0));
        let _436_v36;
        _436_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_395_v6);
        let _437_v37;
        _437_v37 = _dafny.Set.fromElements((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))], (new BigNumber((_436_v36).length)).multipliedBy((_dafny.ZERO).minus((_397_v8)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_397_v8).length))])), new BigNumber(350), p1, new BigNumber(-914));
        let _rhs75 = _dafny.Seq.of(_433_v33, _433_v33, _module.D4.create_DC14(), _433_v33, _433_v33);
        let _rhs76 = _module.__default.fm27(globalState);
        let _rhs77 = _435_v35;
        let _rhs78 = (_437_v37).Difference((_437_v37).Difference(_437_v37));
        let _rhs79 = _dafny.Seq.UnicodeFromString("cps");
        _434_v34 = _rhs75;
        _389_v0 = _rhs76;
        _435_v35 = _rhs77;
        _437_v37 = _rhs78;
        _389_v0 = _rhs79;
      } else {
        let _438_v38;
        _438_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _439_v39;
        _439_v39 = _module.D0.create_DC1(p2);
        let _440_v40;
        _440_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_438_v38).length),_439_v39);
        let _441_v41;
        _441_v41 = _module.D0.create_DC2((((_440_v40).contains(p1)) ? ((_440_v40).get(p1)) : (_439_v39)));
        _441_v41 = _441_v41;
        let _442_v42;
        _442_v42 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber(941)).plus(p1)),(_dafny.ZERO).minus(p1));
        let _443_v43;
        _443_v43 = _dafny.Set.fromElements(p0, p2, p0, p2);
        _442_v42 = (_442_v42).update(new BigNumber((_443_v43).length), new BigNumber((_dafny.Set.fromElements(p2)).length));
        let _444_v45;
        _444_v45 = _dafny.Set.fromElements(p1, new BigNumber(584));
        r2 = ((_444_v45).Union(_444_v45)).IsProperSubsetOf((function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of _dafny.IntegerRange(new BigNumber(508), new BigNumber(914))) {
            let _445_v44 = _compr_23;
            if (((new BigNumber(508)).isLessThanOrEqualTo(_445_v44)) && ((_445_v44).isLessThan(new BigNumber(914)))) {
              _coll23.add((_445_v44).plus((_dafny.ZERO).minus(p1)));
            }
          }
          return _coll23;
        }()).Union(_444_v45));
        r0 = p1;
        let _446_v47;
        _446_v47 = _dafny.Seq.of((function () {
          let _coll24 = new _dafny.Set();
          for (const _compr_24 of _dafny.IntegerRange(new BigNumber(807), new BigNumber(152))) {
            let _447_v46 = _compr_24;
            if (((new BigNumber(807)).isLessThanOrEqualTo(_447_v46)) && ((_447_v46).isLessThan(new BigNumber(152)))) {
              _coll24.add((_447_v46).multipliedBy(new BigNumber((_dafny.Seq.of(p0)).length)));
            }
          }
          return _coll24;
        }()).Difference(_444_v45));
        (globalState).f5 = new BigNumber(((_446_v47)[_module.__default.safeIndex(new BigNumber((_389_v0).length), new BigNumber((_446_v47).length))]).length);
      }
      (globalState).f0 = ((p2) ? (p0) : (p0));
      let _448_v48;
      let _nw78 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _448_v48 = _nw78;
      let _index53 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_448_v48).length));
      (_448_v48)[_index53] = p1;
      let _index54 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_448_v48).length));
      (_448_v48)[_index54] = p1;
      r0 = p1;
      let _449_v49;
      _449_v49 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_448_v48)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_448_v48).length))]);
      let _450_v50;
      _450_v50 = _dafny.Seq.of(new BigNumber(321));
      let _451_v51;
      _451_v51 = _dafny.Map.Empty.slice().updateUnsafe((_450_v50)[_module.__default.safeIndex(p1, new BigNumber((_450_v50).length))],false);
      r0 = new BigNumber((_module.__default.fm2(p0, true, p2, (_449_v49).Merge(_dafny.Map.Empty.slice().updateUnsafe((((_451_v51).contains((_448_v48)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_448_v48).length))])) ? ((_451_v51).get((_448_v48)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_448_v48).length))])) : (p0)),p1)), globalState)).length);
      r1 = p0;
      r2 = !((p2) || (((p0) ? (p2) : (p0))));
      return [r0, r1, r2];
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = [];
      let r3 = _dafny.ZERO;
      let _452_v0;
      _452_v0 = new _dafny.CodePoint('u'.codePointAt(0));
      let _453_v1;
      _453_v1 = _module.D0.create_DC1(p1);
      let _454_v2;
      _454_v2 = _module.D2.create_DC6((_453_v1).dtor_cf1);
      let _455_v3;
      _455_v3 = _dafny.Map.Empty.slice().updateUnsafe(p3,false);
      let _456_v4;
      _456_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_455_v3).contains(p3)) ? ((_455_v3).get(p3)) : (p1)));
      let _457_v5;
      _457_v5 = _dafny.Map.Empty.slice().updateUnsafe(_456_v4,p3);
      let _458_v6;
      _458_v6 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p3));
      let _459_v7;
      let _nw79 = Array((new BigNumber(17)).toNumber());
      _nw79[(_dafny.ZERO).toNumber()] = p0;
      _nw79[(_dafny.ONE).toNumber()] = p1;
      _nw79[(new BigNumber(2)).toNumber()] = p1;
      _nw79[(new BigNumber(3)).toNumber()] = p0;
      _nw79[(new BigNumber(4)).toNumber()] = p0;
      _nw79[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(_452_v0)).IsProperSubsetOf(_dafny.Set.fromElements(_452_v0, _452_v0));
      _nw79[(new BigNumber(6)).toNumber()] = true;
      _nw79[(new BigNumber(7)).toNumber()] = p0;
      _nw79[(new BigNumber(8)).toNumber()] = p1;
      _nw79[(new BigNumber(9)).toNumber()] = (p0) && ((_454_v2).dtor_cf6);
      _nw79[(new BigNumber(10)).toNumber()] = (_457_v5).contains(_456_v4);
      _nw79[(new BigNumber(11)).toNumber()] = false;
      _nw79[(new BigNumber(12)).toNumber()] = ((_458_v6).update(p3, _module.__default.abs(p3))).IsDisjointFrom(_dafny.MultiSet.fromElements(p3));
      _nw79[(new BigNumber(13)).toNumber()] = p1;
      _nw79[(new BigNumber(14)).toNumber()] = p1;
      _nw79[(new BigNumber(15)).toNumber()] = p0;
      _nw79[(new BigNumber(16)).toNumber()] = p0;
      _459_v7 = _nw79;
      let _index55 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length));
      (_459_v7)[_index55] = p0;
      let _460_v8;
      _460_v8 = _module.D3.create_DC9(p3);
      let _pat_let_tv11 = p1;
      let _pat_let_tv12 = p0;
      let _pat_let_tv13 = p1;
      let _index56 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length));
      (_459_v7)[_index56] = function (_source7) {
        if (_source7.is_DC9) {
          let _461___mcc_h0 = (_source7).cf11;
          let _462_cf11 = _461___mcc_h0;
          return _pat_let_tv11;
        } else if (_source7.is_DC8) {
          let _463___mcc_h1 = (_source7).cf10;
          let _464_cf10 = _463___mcc_h1;
          return _pat_let_tv12;
        } else {
          let _465___mcc_h2 = (_source7).cf12;
          let _466_cf12 = _465___mcc_h2;
          return _pat_let_tv13;
        }
      }(_460_v8);
      let _467_v9;
      _467_v9 = _module.D4.create_DC11(_459_v7);
      let _source8 = _467_v9;
      if (_source8.is_DC12) {
        let _468___mcc_h3 = (_source8).cf14;
        let _469___mcc_h4 = (_source8).cf15;
        let _470___mcc_h5 = (_source8).cf16;
        let _471___mcc_h6 = (_source8).cf17;
        let _472___mcc_h7 = (_source8).cf18;
        let _473_cf18 = _472___mcc_h7;
        let _474_cf17 = _471___mcc_h6;
        let _475_cf16 = _470___mcc_h5;
        let _476_cf15 = _469___mcc_h4;
        let _477_cf14 = _468___mcc_h3;
        _475_cf16 = _dafny.Seq.Concat(_dafny.Seq.Concat(_475_cf16, _475_cf16), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("taufebrr"), _475_cf16));
        r3 = _477_cf14;
        let _478_v10;
        _478_v10 = _dafny.Set.fromElements(_452_v0);
        let _479_v13;
        _479_v13 = _dafny.Map.Empty.slice().updateUnsafe(_452_v0,_477_cf14);
        let _480_v15;
        _480_v15 = _dafny.MultiSet.fromElements(_452_v0, _452_v0);
        let _481_v17;
        _481_v17 = _dafny.Map.Empty.slice().updateUnsafe(_452_v0,p1);
        let _482_v19;
        let _nw80 = Array((new BigNumber(22)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = (function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of (function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of (_478_v10).Elements) {
              let _483_v11 = _compr_26;
              if ((_478_v10).contains(_483_v11)) {
                _coll26.add(_483_v11);
              }
            }
            return _coll26;
          }()).Elements) {
            let _484_v12 = _compr_25;
            if ((function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of (_478_v10).Elements) {
                let _485_v11 = _compr_27;
                if ((_478_v10).contains(_485_v11)) {
                  _coll27.add(_485_v11);
                }
              }
              return _coll27;
            }()).contains(_484_v12)) {
              _coll25.add(_484_v12);
            }
          }
          return _coll25;
        }()).Union(_478_v10);
        _nw80[(_dafny.ONE).toNumber()] = _478_v10;
        _nw80[(new BigNumber(2)).toNumber()] = (_478_v10).Intersect(_module.__default.fm28(new BigNumber((_458_v6).cardinality()), _474_cf17, !(p1), globalState));
        _nw80[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_452_v0, _module.__default.fm29(globalState), _452_v0, _452_v0);
        _nw80[(new BigNumber(4)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(_452_v0, _452_v0)).Intersect(_478_v10);
        _nw80[(new BigNumber(6)).toNumber()] = ((true) ? (function () {
          let _coll28 = new _dafny.Set();
          for (const _compr_28 of (_479_v13).Keys.Elements) {
            let _486_v14 = _compr_28;
            if ((_479_v13).contains(_486_v14)) {
              _coll28.add(_486_v14);
            }
          }
          return _coll28;
        }()) : (_dafny.Set.fromElements(_452_v0)));
        _nw80[(new BigNumber(7)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(8)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(9)).toNumber()] = (_478_v10).Intersect(_478_v10);
        _nw80[(new BigNumber(10)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(11)).toNumber()] = (function () {
          let _coll29 = new _dafny.Set();
          for (const _compr_29 of (_480_v15).Elements) {
            let _487_v16 = _compr_29;
            if ((_480_v15).contains(_487_v16)) {
              _coll29.add(_487_v16);
            }
          }
          return _coll29;
        }()).Difference(_478_v10);
        _nw80[(new BigNumber(12)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(13)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(14)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(15)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(16)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(17)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(18)).toNumber()] = (_478_v10).Intersect(function () {
          let _coll30 = new _dafny.Set();
          for (const _compr_30 of (_481_v17).Keys.Elements) {
            let _488_v18 = _compr_30;
            if ((_481_v17).contains(_488_v18)) {
              _coll30.add(_488_v18);
            }
          }
          return _coll30;
        }());
        _nw80[(new BigNumber(19)).toNumber()] = _478_v10;
        _nw80[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements(_452_v0);
        _nw80[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(_452_v0);
        _482_v19 = _nw80;
        let _index57 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_482_v19).length));
        (_482_v19)[_index57] = _module.__default.fm28(_474_cf17, new BigNumber((_475_cf16).length), (_459_v7)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length))], globalState);
        let _index58 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_482_v19).length));
        (_482_v19)[_index58] = _478_v10;
        r3 = new BigNumber(340);
      } else if (_source8.is_DC13) {
        let _489___mcc_h8 = (_source8).cf19;
        let _490___mcc_h9 = (_source8).cf20;
        let _491_cf20 = _490___mcc_h9;
        let _492_cf19 = _489___mcc_h8;
        let _493_v20;
        let _nw81 = new _module.C0();
        _nw81.__ctor((p2).IsSubsetOf(p2), _459_v7);
        _493_v20 = _nw81;
        let _494_v21;
        let _nw82 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _494_v21 = _nw82;
        _494_v21 = _494_v21;
        _492_cf19 = _492_cf19;
        let _index59 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_494_v21).length));
        (_494_v21)[_index59] = p3;
        let _index60 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_494_v21).length));
        (_494_v21)[_index60] = _492_cf19;
      } else if (_source8.is_DC14) {
        let _495_v22;
        _495_v22 = _dafny.Seq.UnicodeFromString("guejok");
        let _496_v23;
        let _nw83 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _496_v23 = _nw83;
        let _497_v24;
        _497_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,_496_v23);
        let _498_v25;
        _498_v25 = _module.D4.create_DC12(_module.__default.fm0(new BigNumber((_dafny.Seq.UnicodeFromString("isuftd")).length), p3, p3, true, globalState), p3, _495_v22, (_dafny.ZERO).minus(new BigNumber(((_497_v24).update(true, _496_v23)).length)), new BigNumber(612));
        r3 = new BigNumber(((_498_v25).dtor_cf16).length);
        let _499_v26;
        let _init14 = ((_500_v0) => function (_501_i0) {
          return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-92)), ((_502_v0) => function (_503_i1) {
            return _502_v0;
          })(_500_v0)), _dafny.Seq.UnicodeFromString("elvverckl"));
        })(_452_v0);
        let _nw84 = Array((new BigNumber(24)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw84.length); _i0_14++) {
          _nw84[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _499_v26 = _nw84;
        let _504_v27;
        _504_v27 = _dafny.Seq.of(_495_v22, _495_v22, _495_v22);
        let _505_v28;
        _505_v28 = _dafny.Seq.of(p0, p0);
        let _506_v29;
        _506_v29 = _module.D8.create_DC21(p0, p1, _505_v28, true, p3);
        let _507_v30;
        _507_v30 = _dafny.Set.fromElements((_505_v28)[_module.__default.safeIndex(new BigNumber(((_506_v29).dtor_cf28).length), new BigNumber((_505_v28).length))]);
        let _index61 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_499_v26).length));
        (_499_v26)[_index61] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(638)), function (_508_i2) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), (_504_v27)[_module.__default.safeIndex(new BigNumber((_507_v30).length), new BigNumber((_504_v27).length))]);
        let _index62 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_499_v26).length));
        (_499_v26)[_index62] = _495_v22;
        let _509_v31;
        _509_v31 = _module.D0.create_DC0(_455_v3);
        let _510_v32;
        _510_v32 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, p0)));
        let _511_v33;
        let _nw85 = new _module.C0();
        _nw85.__ctor((_this).fm4(_509_v31, _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(728),(_459_v7)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length))])), (_510_v32)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_510_v32).length))], globalState), _459_v7);
        _511_v33 = _nw85;
        let _512_v34;
        let _nw86 = new _module.C0();
        _nw86.__ctor((_454_v2).dtor_cf6, (_511_v33).f11);
        _512_v34 = _nw86;
      } else {
        let _513___mcc_h10 = (_source8).cf13;
        let _514_cf13 = _513___mcc_h10;
        (globalState).f5 = p3;
        let _515_v35;
        let _nw87 = new _module.C0();
        _nw87.__ctor(p0, _514_cf13);
        _515_v35 = _nw87;
        let _516_v36;
        let _init15 = ((_517_p3) => function (_518_i3) {
          return (_518_i3).minus(_517_p3);
        })(p3);
        let _nw88 = Array((new BigNumber(12)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw88.length); _i0_15++) {
          _nw88[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _516_v36 = _nw88;
        let _519_v37;
        _519_v37 = _dafny.Seq.of(_516_v36);
        let _520_v38;
        _520_v38 = _module.D10.create_DC26(_519_v37);
        _519_v37 = _dafny.Seq.Concat((_520_v38).dtor_cf37, _519_v37);
        let _nw89 = new _module.C0();
        _nw89.__ctor((p3).isLessThanOrEqualTo(p3), (_515_v35).f11);
        _515_v35 = _nw89;
      }
      r3 = _module.__default.safeDivisionInt(p3, p3);
      let _521_v39;
      let _nw90 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _521_v39 = _nw90;
      let _522_v40;
      _522_v40 = _521_v39;
      let _source9 = _522_v40;
      let _523___mcc_h11 = _source9;
      let _524_cf21 = _523___mcc_h11;
      let _525_v41;
      _525_v41 = _dafny.MultiSet.fromElements(_452_v0);
      _525_v41 = ((_525_v41).Union(_525_v41)).Union((_525_v41).update(_452_v0, _module.__default.abs(p3)));
      (globalState).f5 = p3;
      let _526_v43;
      _526_v43 = _dafny.Seq.of(p3);
      let _527_v44;
      _527_v44 = _dafny.Seq.of((_455_v3).Merge(function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of (_526_v43).Elements) {
          let _528_v42 = _compr_31;
          if (_dafny.Seq.contains(_526_v43, _528_v42)) {
            _coll31.push([_module.__default.safeDivisionInt(_528_v42, new BigNumber(-970)),p1]);
          }
        }
        return _coll31;
      }()), _dafny.Map.Empty.slice().updateUnsafe(p3,(_459_v7)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length))]), _455_v3);
      _527_v44 = _dafny.Seq.update(_dafny.Seq.Concat(_527_v44, _527_v44), _module.__default.safeIndex(_module.__default.fm0(p3, p3, p3, (_459_v7)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length))], globalState), new BigNumber((_dafny.Seq.Concat(_527_v44, _527_v44)).length)), _module.__default.fm22(p0, globalState));
      let _529_v45;
      let _init16 = function (_530_i4) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ipgn"), _dafny.Seq.UnicodeFromString("uhyhwmg"));
      };
      let _nw91 = Array((new BigNumber(9)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw91.length); _i0_16++) {
        _nw91[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _529_v45 = _nw91;
      let _531_v46;
      _531_v46 = _dafny.Seq.UnicodeFromString("knmngv");
      let _index63 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_529_v45).length));
      (_529_v45)[_index63] = _531_v46;
      let _index64 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_529_v45).length));
      (_529_v45)[_index64] = _module.__default.fm27(globalState);
      let _532_v47;
      let _nw92 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _532_v47 = _nw92;
      let _index65 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_532_v47).length));
      (_532_v47)[_index65] = _452_v0;
      let _index66 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_532_v47).length));
      (_532_v47)[_index66] = _452_v0;
      let _533_v48;
      _533_v48 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(-115));
      let _534_v49;
      _534_v49 = _module.D0.create_DC0(_455_v3);
      let _rhs80 = ((p0) ? (_533_v48) : (_533_v48));
      let _rhs81 = (((_456_v4).contains(!((p1) === (p1)))) ? ((_456_v4).get(!((p1) === (p1)))) : ((_this).fm4(_534_v49, _534_v49, p2, globalState)));
      let _rhs82 = _459_v7;
      _533_v48 = _rhs80;
      r1 = _rhs81;
      _459_v7 = _rhs82;
      let _535_v50;
      _535_v50 = _dafny.Map.Empty.slice().updateUnsafe((_459_v7)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length))],new BigNumber(-467));
      r0 = (_535_v50).update(p0, p3);
      r1 = (_459_v7)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_459_v7).length))];
      let _536_v51;
      let _nw93 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
      _536_v51 = _nw93;
      let _537_v52;
      _537_v52 = _dafny.Seq.of(_536_v51, _536_v51, _536_v51);
      r2 = (_537_v52)[_module.__default.safeIndex(p3, new BigNumber((_537_v52).length))];
      r3 = p3;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f13 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f13) {
      let _this = this;
      (_this)._f13 = f13;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f13;
    };
    fm18(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(65);
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _538_v0;
      _538_v0 = _dafny.Seq.of((_this).f13, (_this).f13);
      r2 = !(_dafny.areEqual(_dafny.Seq.Concat(_538_v0, _538_v0), _dafny.Seq.Concat(_538_v0, _538_v0)));
      let _539_v1;
      _539_v1 = new _dafny.CodePoint('j'.codePointAt(0));
      let _540_v2;
      _540_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,_539_v1);
      let _541_v3;
      _541_v3 = _dafny.Seq.UnicodeFromString("bnmkw");
      let _542_v4;
      _542_v4 = _module.D3.create_DC8(_541_v3);
      let _543_v5;
      _543_v5 = _module.D3.create_DC10(_542_v4);
      let _544_v7;
      _544_v7 = _dafny.Map.Empty.slice().updateUnsafe(_543_v5,function () {
        let _coll32 = new _dafny.Map();
        for (const _compr_32 of _dafny.IntegerRange(new BigNumber(225), new BigNumber(316))) {
          let _545_v6 = _compr_32;
          if (((new BigNumber(225)).isLessThanOrEqualTo(_545_v6)) && ((_545_v6).isLessThan(new BigNumber(316)))) {
            _coll32.push([(_545_v6).multipliedBy(p2),_539_v1]);
          }
        }
        return _coll32;
      }());
      let _546_v9;
      _546_v9 = _dafny.Set.fromElements(p0, p2);
      let _547_v10;
      _547_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of (_546_v9).Elements) {
          let _548_v8 = _compr_33;
          if ((_546_v9).contains(_548_v8)) {
            _coll33.push([_module.__default.safeDivisionInt(_548_v8, p2),_539_v1]);
          }
        }
        return _coll33;
      }());
      let _549_v11;
      _549_v11 = _dafny.Set.fromElements((_this).f13, (_this).f13);
      let _550_v12;
      _550_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber((_549_v11).length));
      let _551_v14;
      _551_v14 = _dafny.Seq.of(_540_v2);
      let _552_v15;
      _552_v15 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(new BigNumber((_541_v3).length)));
      let _553_v17;
      let _nw94 = Array((new BigNumber(28)).toNumber());
      _nw94[(_dafny.ZERO).toNumber()] = (_540_v2).update(p2, _539_v1);
      _nw94[(_dafny.ONE).toNumber()] = (_540_v2).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(952),_539_v1));
      _nw94[(new BigNumber(2)).toNumber()] = ((_540_v2).update(p0, _539_v1)).Merge(_540_v2);
      _nw94[(new BigNumber(3)).toNumber()] = (_540_v2).Merge(_540_v2);
      _nw94[(new BigNumber(4)).toNumber()] = ((((_544_v7).contains(_543_v5)) ? ((_544_v7).get(_543_v5)) : (_540_v2))).Merge(_540_v2);
      _nw94[(new BigNumber(5)).toNumber()] = _module.__default.fm19(globalState);
      _nw94[(new BigNumber(6)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(7)).toNumber()] = (_540_v2).Merge(_540_v2);
      _nw94[(new BigNumber(8)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(9)).toNumber()] = (_540_v2).update(p2, new _dafny.CodePoint('d'.codePointAt(0)));
      _nw94[(new BigNumber(10)).toNumber()] = (((_this).f13) ? (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),_539_v1)) : (_540_v2));
      _nw94[(new BigNumber(11)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p2,(_541_v3)[_module.__default.safeIndex(p0, new BigNumber((_541_v3).length))]);
      _nw94[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13)).length),_539_v1)).update(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(617)), ((_554_v1) => function (_555_i1) {
        return _554_v1;
      })(_539_v1))).length), _539_v1);
      _nw94[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p2,_539_v1);
      _nw94[(new BigNumber(15)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(16)).toNumber()] = (_540_v2).update(p3, _539_v1);
      _nw94[(new BigNumber(17)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(18)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(19)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(20)).toNumber()] = (((_547_v10).contains(!((_this).f13))) ? ((_547_v10).get(!((_this).f13))) : (_540_v2));
      _nw94[(new BigNumber(21)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(22)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(23)).toNumber()] = _540_v2;
      _nw94[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm20((_dafny.ZERO).minus(new BigNumber((_550_v12).length)), (_this).f13, p3, globalState));
      _nw94[(new BigNumber(25)).toNumber()] = function () {
        let _coll34 = new _dafny.Map();
        for (const _compr_34 of (_546_v9).Elements) {
          let _556_v13 = _compr_34;
          if ((_546_v9).contains(_556_v13)) {
            _coll34.push([_module.__default.safeDivisionInt(_556_v13, p3),_539_v1]);
          }
        }
        return _coll34;
      }();
      _nw94[(new BigNumber(26)).toNumber()] = (_551_v14)[_module.__default.safeIndex((_552_v15)[_module.__default.safeIndex(new BigNumber((_541_v3).length), new BigNumber((_552_v15).length))], new BigNumber((_551_v14).length))];
      _nw94[(new BigNumber(27)).toNumber()] = (function () {
        let _coll35 = new _dafny.Map();
        for (const _compr_35 of _dafny.IntegerRange(new BigNumber(921), new BigNumber(-898))) {
          let _557_v16 = _compr_35;
          if (((new BigNumber(921)).isLessThanOrEqualTo(_557_v16)) && ((_557_v16).isLessThan(new BigNumber(-898)))) {
            _coll35.push([(_557_v16).multipliedBy(p2),new _dafny.CodePoint('s'.codePointAt(0))]);
          }
        }
        return _coll35;
      }()).update(p2, _539_v1);
      _553_v17 = _nw94;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_553_v17).length))) {
        let _558_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_558_i0)) && ((_558_i0).isLessThan(new BigNumber((_553_v17).length))))) {
          (_553_v17)[(_558_i0)] = (_540_v2).Merge((_540_v2).update(new BigNumber(74), _539_v1));
        }
      }
      let _559_i2;
      _559_i2 = _dafny.ZERO;
      L3: {
        while ((_this).f13) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_559_i2)) {
              break L3;
            }
            _559_i2 = (_559_i2).plus(_dafny.ONE);
            (globalState).f5 = (_dafny.ZERO).minus((new BigNumber(65)).multipliedBy(p0));
            r1 = (_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f13)).length), p2)).minus((new BigNumber((_dafny.Seq.UnicodeFromString("mow")).length)).plus((_dafny.ZERO).minus(p3)));
            if ((_this).f13) {
              let _560_v18;
              let _init17 = ((_561_p2) => function (_562_i3) {
                return (_562_i3).plus(_561_p2);
              })(p2);
              let _nw95 = Array((new BigNumber(21)).toNumber());
              for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw95.length); _i0_17++) {
                _nw95[_i0_17] = _init17(new BigNumber(_i0_17));
              }
              _560_v18 = _nw95;
              let _index67 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v18).length));
              (_560_v18)[_index67] = (p0).plus(p3);
              let _index68 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v18).length));
              (_560_v18)[_index68] = p0;
              let _563_v19;
              let _nw96 = Array((new BigNumber(7)).toNumber());
              _nw96[(_dafny.ZERO).toNumber()] = (_this).f13;
              _nw96[(_dafny.ONE).toNumber()] = false;
              _nw96[(new BigNumber(2)).toNumber()] = (((_this).f13) ? ((_538_v0)[_module.__default.safeIndex(new BigNumber((_541_v3).length), new BigNumber((_538_v0).length))]) : ((_this).f13));
              _nw96[(new BigNumber(3)).toNumber()] = (_this).f13;
              _nw96[(new BigNumber(4)).toNumber()] = false;
              _nw96[(new BigNumber(5)).toNumber()] = (_this).f13;
              _nw96[(new BigNumber(6)).toNumber()] = (_this).f13;
              _563_v19 = _nw96;
              _563_v19 = _563_v19;
              let _564_v20;
              let _nw97 = new _module.C0();
              _nw97.__ctor(true, _563_v19);
              _564_v20 = _nw97;
              let _565_v21;
              _565_v21 = _dafny.Set.fromElements(_module.D2.create_DC5(_564_v20));
              let _566_v22;
              _566_v22 = _module.D2.create_DC5(_564_v20);
              let _pat_let_tv14 = _564_v20;
              _565_v21 = (((_this).f13) ? (_dafny.Set.fromElements(function (_pat_let6_0) {
                return function (_567_dt__update__tmp_h0) {
                  return function (_pat_let7_0) {
                    return function (_568_dt__update_hcf5_h0) {
                      return _module.D2.create_DC5(_568_dt__update_hcf5_h0);
                    }(_pat_let7_0);
                  }(_pat_let_tv14);
                }(_pat_let6_0);
              }(_module.D2.create_DC5(_564_v20)), _566_v22)) : (_dafny.Set.fromElements(_module.D2.create_DC5(_564_v20), _566_v22, _566_v22, _566_v22, _566_v22)));
              _550_v12 = (_550_v12).update(true, _module.__default.safeDivisionInt((_560_v18)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v18).length))], p2));
              let _569_v23;
              _569_v23 = _module.D2.create_DC6((_this).f13);
              let _index69 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v18).length));
              let _rhs83 = (p3).multipliedBy((_this).fm18(_569_v23, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm21(globalState)).length)), p3, globalState));
              let _rhs84 = (_this).f13;
              let _rhs85 = _560_v18;
              let _lhs58 = _560_v18;
              let _lhs59 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v18).length));
              let _lhs60 = globalState;
              _lhs58[_lhs59] = _rhs83;
              _lhs60.f2 = _rhs84;
              _560_v18 = _rhs85;
            } else {
              let _570_v24;
              _570_v24 = _module.D2.create_DC6((_this).f13);
              let _571_v25;
              let _nw98 = Array((new BigNumber(13)).toNumber());
              _nw98[(_dafny.ZERO).toNumber()] = (_this).f13;
              _nw98[(_dafny.ONE).toNumber()] = (_this).f13;
              _nw98[(new BigNumber(2)).toNumber()] = (_module.__default.fm1((_this).f13, globalState)) || ((_this).f13);
              _nw98[(new BigNumber(3)).toNumber()] = (((_this).f13) ? ((_this).f13) : ((_this).f13));
              _nw98[(new BigNumber(4)).toNumber()] = false;
              _nw98[(new BigNumber(5)).toNumber()] = (_this).f13;
              _nw98[(new BigNumber(6)).toNumber()] = (_this).f13;
              _nw98[(new BigNumber(7)).toNumber()] = (_this).f13;
              _nw98[(new BigNumber(8)).toNumber()] = (_this).f13;
              _nw98[(new BigNumber(9)).toNumber()] = (_this).f13;
              _nw98[(new BigNumber(10)).toNumber()] = (_570_v24).dtor_cf6;
              _nw98[(new BigNumber(11)).toNumber()] = (_this).f13;
              _nw98[(new BigNumber(12)).toNumber()] = (_this).f13;
              _571_v25 = _nw98;
              let _index70 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_571_v25).length));
              (_571_v25)[_index70] = (_this).f13;
              let _index71 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_571_v25).length));
              (_571_v25)[_index71] = !(!(p3).isEqualTo(p3));
              let _572_v26;
              let _init18 = ((_573_v1) => function (_574_i4) {
                return _573_v1;
              })(_539_v1);
              let _nw99 = Array((new BigNumber(7)).toNumber());
              for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw99.length); _i0_18++) {
                _nw99[_i0_18] = _init18(new BigNumber(_i0_18));
              }
              _572_v26 = _nw99;
              let _index72 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_572_v26).length));
              (_572_v26)[_index72] = (((_this).f13) ? (_539_v1) : (_539_v1));
              let _index73 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_571_v25).length));
              let _index74 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_572_v26).length));
              let _rhs86 = (_571_v25)[_module.__default.safeIndex(new BigNumber(295), new BigNumber((_571_v25).length))];
              let _rhs87 = _539_v1;
              let _lhs61 = _571_v25;
              let _lhs62 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_571_v25).length));
              let _lhs63 = _572_v26;
              let _lhs64 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_572_v26).length));
              _lhs61[_lhs62] = _rhs86;
              _lhs63[_lhs64] = _rhs87;
              (globalState).f2 = (_this).f13;
              (globalState).f5 = p3;
              let _575_v27;
              _575_v27 = _module.D4.create_DC12(p0, p3, _541_v3, p3, new BigNumber((_538_v0).length));
              let _576_v28;
              _576_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_575_v27);
              _576_v28 = (_576_v28).update(p2, _575_v27);
            }
            let _577_v29;
            let _nw100 = new _module.C1();
            _nw100.__ctor();
            _577_v29 = _nw100;
            let _578_v30;
            let _nw101 = Array((new BigNumber(7)).toNumber());
            _nw101[(_dafny.ZERO).toNumber()] = _577_v29;
            _nw101[(_dafny.ONE).toNumber()] = _577_v29;
            _nw101[(new BigNumber(2)).toNumber()] = _577_v29;
            _nw101[(new BigNumber(3)).toNumber()] = _577_v29;
            _nw101[(new BigNumber(4)).toNumber()] = _577_v29;
            _nw101[(new BigNumber(5)).toNumber()] = _577_v29;
            _nw101[(new BigNumber(6)).toNumber()] = _577_v29;
            _578_v30 = _nw101;
            let _579_v31;
            _579_v31 = _module.D11.create_DC29(_577_v29);
            let _index75 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_578_v30).length));
            (_578_v30)[_index75] = (_579_v31).dtor_cf43;
            let _index76 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_578_v30).length));
            (_578_v30)[_index76] = _577_v29;
          }
        }
      }
      let _580_v32;
      let _nw102 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _580_v32 = _nw102;
      _580_v32 = _580_v32;
      let _581_v33;
      let _nw103 = new _module.C1();
      _nw103.__ctor();
      _581_v33 = _nw103;
      _581_v33 = _581_v33;
      let _582_v34;
      _582_v34 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f13);
      let _583_v36;
      _583_v36 = _module.D2.create_DC6((_this).f13);
      let _584_v38;
      let _nw104 = Array((new BigNumber(17)).toNumber());
      _nw104[(_dafny.ZERO).toNumber()] = _582_v34;
      _nw104[(_dafny.ONE).toNumber()] = (function () {
        let _coll36 = new _dafny.Map();
        for (const _compr_36 of _dafny.IntegerRange(new BigNumber(-111), new BigNumber(985))) {
          let _585_v35 = _compr_36;
          if (((new BigNumber(-111)).isLessThanOrEqualTo(_585_v35)) && ((_585_v35).isLessThan(new BigNumber(985)))) {
            _coll36.push([(_585_v35).minus(p0),false]);
          }
        }
        return _coll36;
      }()).Merge(_582_v34);
      _nw104[(new BigNumber(2)).toNumber()] = (_582_v34).Merge(_582_v34);
      _nw104[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).fm18(_583_v36, new BigNumber(510), p2, globalState),(_this).f13);
      _nw104[(new BigNumber(4)).toNumber()] = (_582_v34).update((_dafny.ZERO).minus(p0), (_this).f13);
      _nw104[(new BigNumber(5)).toNumber()] = (_582_v34).Merge(function () {
        let _coll37 = new _dafny.Map();
        for (const _compr_37 of _dafny.IntegerRange(new BigNumber(278), new BigNumber(751))) {
          let _586_v37 = _compr_37;
          if (((new BigNumber(278)).isLessThanOrEqualTo(_586_v37)) && ((_586_v37).isLessThan(new BigNumber(751)))) {
            _coll37.push([(_586_v37).multipliedBy(p0),(_this).f13]);
          }
        }
        return _coll37;
      }());
      _nw104[(new BigNumber(6)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(7)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(8)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(9)).toNumber()] = (_582_v34).update(p0, (_this).f13);
      _nw104[(new BigNumber(10)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(11)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(12)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(13)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(14)).toNumber()] = _582_v34;
      _nw104[(new BigNumber(15)).toNumber()] = (_582_v34).Merge(_582_v34);
      _nw104[(new BigNumber(16)).toNumber()] = (_582_v34).Merge(_582_v34);
      _584_v38 = _nw104;
      let _587_v39;
      _587_v39 = _dafny.Set.fromElements((_541_v3)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_541_v3).length))], new _dafny.CodePoint('j'.codePointAt(0)));
      let _588_v40;
      _588_v40 = _dafny.Map.Empty.slice().updateUnsafe((_587_v39).Difference(_587_v39),_584_v38);
      let _rhs88 = (_this).f13;
      let _rhs89 = _541_v3;
      let _rhs90 = p3;
      let _rhs91 = (((_588_v40).contains(_587_v39)) ? ((_588_v40).get(_587_v39)) : (_584_v38));
      let _lhs65 = globalState;
      r3 = _rhs88;
      _541_v3 = _rhs89;
      _lhs65.f5 = _rhs90;
      _584_v38 = _rhs91;
      let _589_v41;
      _589_v41 = _dafny.Map.Empty.slice().updateUnsafe(p2,_538_v0);
      r0 = (new BigNumber(757)).isLessThan(new BigNumber((_589_v41).length));
      r1 = _module.__default.safeDivisionInt(p0, p0);
      r2 = !(!(false));
      let _590_v42;
      _590_v42 = _module.D0.create_DC0(_582_v34);
      let _591_v43;
      _591_v43 = _dafny.MultiSet.fromElements((_this).f13);
      let _592_v44;
      _592_v44 = _module.D8.create_DC21((_this).f13, (_this).f13, _538_v0, (_581_v33).fm4(_590_v42, _590_v42, _591_v43, globalState), p0);
      r3 = !((_592_v44).dtor_cf29);
      return [r0, r1, r2, r3];
    }
    m4(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.MultiSet.Empty;
      let r2 = _dafny.Map.Empty;
      let _593_v0;
      _593_v0 = new BigNumber(945);
      let _hi2 = _593_v0;
      for (let _594_i0 = _593_v0; _594_i0.isLessThan(_hi2); _594_i0 = _594_i0.plus(_dafny.ONE)) {
        _593_v0 = _594_i0;
        let _595_v1;
        let _nw105 = new _module.C1();
        _nw105.__ctor();
        _595_v1 = _nw105;
        let _596_v2;
        let _init19 = function (_597_i1) {
          return (_this).f13;
        };
        let _nw106 = Array((new BigNumber(14)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw106.length); _i0_19++) {
          _nw106[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _596_v2 = _nw106;
        let _598_v3;
        let _nw107 = new _module.C0();
        _nw107.__ctor((_this).f13, _596_v2);
        _598_v3 = _nw107;
        let _599_v4;
        _599_v4 = _module.D2.create_DC5(_598_v3);
        let _600_v5;
        _600_v5 = _dafny.Set.fromElements(_599_v4, _599_v4);
        if ((_600_v5).IsProperSubsetOf((_600_v5).Union(_600_v5))) {
          let _601_v6;
          _601_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_594_i0);
          let _602_v7;
          _602_v7 = _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_601_v6).length),!((_this).f13)));
          (globalState).f2 = (((_this).f13) ? ((_this).f13) : ((_this).fm4(_602_v7, _602_v7, _dafny.MultiSet.fromElements(false), globalState)));
          let _603_v8;
          _603_v8 = _dafny.Set.fromElements(_594_i0, _593_v0, _594_i0, _593_v0);
          let _604_v9;
          _604_v9 = _dafny.Seq.of((new BigNumber((_603_v8).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_module.__default.fm0(new BigNumber((_dafny.MultiSet.fromElements(_594_i0, _593_v0)).cardinality()), _593_v0, _593_v0, (_this).f13, globalState))), (_this).f13, true, (_this).f13, (_this).f13);
          _604_v9 = _dafny.Seq.Concat(_604_v9, _604_v9);
          let _index77 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_596_v2).length));
          (_596_v2)[_index77] = (_this).f13;
          let _index78 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_596_v2).length));
          (_596_v2)[_index78] = (_this).f13;
          _603_v8 = function () {
            let _coll38 = new _dafny.Set();
            for (const _compr_38 of _dafny.IntegerRange(new BigNumber(676), new BigNumber(277))) {
              let _605_v10 = _compr_38;
              if (((new BigNumber(676)).isLessThanOrEqualTo(_605_v10)) && ((_605_v10).isLessThan(new BigNumber(277)))) {
                _coll38.add((_605_v10).multipliedBy((_module.D2.create_DC7(_593_v0, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),(_596_v2)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_596_v2).length))]),_594_i0), _593_v0)).dtor_cf7));
              }
            }
            return _coll38;
          }();
          let _606_v11;
          let _nw108 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _606_v11 = _nw108;
          let _index79 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_606_v11).length));
          (_606_v11)[_index79] = _593_v0;
          let _607_v12;
          _607_v12 = _dafny.Map.Empty.slice().updateUnsafe((_596_v2)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_596_v2).length))],_dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), function (_608_i2) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }));
          let _609_v13;
          _609_v13 = _dafny.Seq.UnicodeFromString("d");
          let _610_v14;
          _610_v14 = _dafny.Seq.of(_609_v13);
          let _index80 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_596_v2).length));
          let _index81 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_606_v11).length));
          let _rhs92 = (((_this).f13) ? ((_593_v0).multipliedBy(_593_v0)) : (_module.__default.safeModuloInt(_593_v0, new BigNumber(((((_607_v12).contains(false)) ? ((_607_v12).get(false)) : (_dafny.Seq.UnicodeFromString("smmjmydnm")))).length))));
          let _rhs93 = !_dafny.areEqual(_610_v14, _610_v14);
          let _rhs94 = _594_i0;
          let _lhs66 = globalState;
          let _lhs67 = _596_v2;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_596_v2).length));
          let _lhs69 = _606_v11;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_606_v11).length));
          _lhs66.f5 = _rhs92;
          _lhs67[_lhs68] = _rhs93;
          _lhs69[_lhs70] = _rhs94;
        } else {
          let _611_v15;
          let _init20 = function (_612_i3) {
            return _dafny.Seq.UnicodeFromString("m");
          };
          let _nw109 = Array((new BigNumber(25)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw109.length); _i0_20++) {
            _nw109[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _611_v15 = _nw109;
          let _index82 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_611_v15).length));
          (_611_v15)[_index82] = _dafny.Seq.UnicodeFromString("fyghxaonv");
          let _613_v16;
          _613_v16 = _dafny.Seq.UnicodeFromString("trl");
          let _index83 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_611_v15).length));
          (_611_v15)[_index83] = _613_v16;
          let _614_v17;
          _614_v17 = _dafny.Seq.of((_this).f13, (_this).f13, (_this).f13, false, (_this).f13);
          let _615_v18;
          _615_v18 = _dafny.MultiSet.fromElements((_this).f13, true);
          r0 = ((((_this).f13) ? (_615_v18) : (_615_v18))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_614_v17));
          let _616_v19;
          _616_v19 = new _dafny.CodePoint('w'.codePointAt(0));
          let _617_v20;
          _617_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_616_v19);
          let _618_v21;
          _618_v21 = _dafny.Set.fromElements((_this).f13);
          let _619_v22;
          _619_v22 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_617_v20).length)).isEqualTo((_dafny.ZERO).minus(_594_i0)),_618_v21);
          _619_v22 = _619_v22;
          let _620_v23;
          let _nw110 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _620_v23 = _nw110;
          let _621_v24;
          _621_v24 = _module.D10.create_DC26(_dafny.Seq.of(_620_v23, _620_v23));
          _621_v24 = _621_v24;
          _593_v0 = _593_v0;
        }
        let _622_v25;
        let _nw111 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
        _622_v25 = _nw111;
        let _623_v26;
        _623_v26 = _dafny.MultiSet.fromElements((_this).f13);
        let _624_v27;
        _624_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_594_i0);
        let _625_v28;
        _625_v28 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index84 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_622_v25).length));
        (_622_v25)[_index84] = (_623_v26).Intersect((_module.D10.create_DC27(true, (((_624_v27).contains((_this).f13)) ? ((_624_v27).get((_this).f13)) : (_594_i0)), _625_v28, _623_v26)).dtor_cf41);
        let _626_v29;
        _626_v29 = _dafny.Set.fromElements(!((_this).f13));
        let _627_v30;
        _627_v30 = _module.D12.create_DC33(_626_v29);
        let _index85 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_622_v25).length));
        (_622_v25)[_index85] = _dafny.MultiSet.fromElements((_this).f13, true, true, ((_627_v30).dtor_cf47).IsSubsetOf(_626_v29));
      }
      let _628_v31;
      let _nw112 = Array((new BigNumber(6)).toNumber()).fill(_module.D2.Default());
      _628_v31 = _nw112;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_628_v31).length))) {
        let _629_i4 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_629_i4)) && ((_629_i4).isLessThan(new BigNumber((_628_v31).length))))) {
          (_628_v31)[(_629_i4)] = _module.D2.create_DC7((_dafny.ZERO).minus(_593_v0), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13),_593_v0), (((_this).f13) ? (new BigNumber((_dafny.Seq.UnicodeFromString("wlpnjgxrs")).length)) : (_593_v0)));
        }
      }
      let _rhs95 = _593_v0;
      let _rhs96 = _module.__default.safeDivisionInt(new BigNumber(177), _593_v0);
      let _lhs71 = globalState;
      _593_v0 = _rhs95;
      _lhs71.f5 = _rhs96;
      let _630_v32;
      let _nw113 = new _module.C1();
      _nw113.__ctor();
      _630_v32 = _nw113;
      let _hi3 = _593_v0;
      for (let _631_i5 = _module.__default.safeDivisionInt(_593_v0, _593_v0); _631_i5.isLessThan(_hi3); _631_i5 = _631_i5.plus(_dafny.ONE)) {
        let _632_v33;
        _632_v33 = _dafny.Map.Empty.slice().updateUnsafe(_593_v0,(_this).f13);
        let _633_v34;
        _633_v34 = _module.D0.create_DC0(_632_v33);
        let _634_v35;
        _634_v35 = _dafny.MultiSet.fromElements((_this).f13, false);
        let _635_v36;
        let _nw114 = Array((new BigNumber(12)).toNumber());
        _nw114[(_dafny.ZERO).toNumber()] = (_this).f13;
        _nw114[(_dafny.ONE).toNumber()] = true;
        _nw114[(new BigNumber(2)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(3)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(4)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(5)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(6)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(7)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(8)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(9)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(10)).toNumber()] = (_this).f13;
        _nw114[(new BigNumber(11)).toNumber()] = false;
        _635_v36 = _nw114;
        let _636_v37;
        _636_v37 = _dafny.Seq.of(_635_v36);
        let _637_v38;
        _637_v38 = _dafny.Set.fromElements(_632_v33, _632_v33);
        let _638_v39;
        _638_v39 = _dafny.Seq.UnicodeFromString("vgk");
        let _639_v40;
        _639_v40 = _dafny.Seq.of((_this).f13, (_this).f13);
        let _640_v41;
        _640_v41 = _module.D12.create_DC36((_this).f13, new BigNumber(75), (_this).f13);
        let _641_v42;
        let _nw115 = new _module.C0();
        _nw115.__ctor((_this).f13, _635_v36);
        _641_v42 = _nw115;
        let _642_v43;
        _642_v43 = _dafny.Seq.of(_641_v42);
        let _643_v44;
        let _nw116 = Array((new BigNumber(29)).toNumber());
        _nw116[(_dafny.ZERO).toNumber()] = _module.__default.fm1((_this).f13, globalState);
        _nw116[(_dafny.ONE).toNumber()] = false;
        _nw116[(new BigNumber(2)).toNumber()] = !((_this).f13);
        _nw116[(new BigNumber(3)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(4)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(5)).toNumber()] = (_this).fm4(_633_v34, _module.D0.create_DC0(_632_v33), _634_v35, globalState);
        _nw116[(new BigNumber(6)).toNumber()] = !_dafny.areEqual(_636_v37, _636_v37);
        _nw116[(new BigNumber(7)).toNumber()] = (_637_v38).IsProperSubsetOf(_637_v38);
        _nw116[(new BigNumber(8)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(9)).toNumber()] = !(_dafny.Seq.IsPrefixOf(_638_v39, _638_v39));
        _nw116[(new BigNumber(10)).toNumber()] = (_634_v35).IsProperSubsetOf(_634_v35);
        _nw116[(new BigNumber(11)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(12)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(13)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(14)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(15)).toNumber()] = (new BigNumber((_dafny.Seq.update(_639_v40, _module.__default.safeIndex(_593_v0, new BigNumber((_639_v40).length)), !((_this).f13))).length)).isLessThan(_631_i5);
        _nw116[(new BigNumber(16)).toNumber()] = true;
        _nw116[(new BigNumber(17)).toNumber()] = !((!(!((_this).f13))) || ((_640_v41).dtor_cf53));
        _nw116[(new BigNumber(18)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(19)).toNumber()] = _dafny.areEqual(_module.__default.fm21(globalState), _638_v39);
        _nw116[(new BigNumber(20)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(21)).toNumber()] = _module.__default.fm1((_this).f13, globalState);
        _nw116[(new BigNumber(22)).toNumber()] = !((_this).f13);
        _nw116[(new BigNumber(23)).toNumber()] = !((_this).f13);
        _nw116[(new BigNumber(24)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(25)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(26)).toNumber()] = (_this).f13;
        _nw116[(new BigNumber(27)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_642_v43, _642_v43);
        _nw116[(new BigNumber(28)).toNumber()] = ((true) ? ((_this).f13) : ((_641_v42).f10));
        _643_v44 = _nw116;
        _643_v44 = (_641_v42).f11;
        let _644_v45;
        let _nw117 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
        _644_v45 = _nw117;
        _644_v45 = _644_v45;
        _641_v42 = _641_v42;
        (globalState).f5 = _module.__default.safeDivisionInt(_631_i5, _631_i5);
      }
      if ((_this).f13) {
        let _645_v46;
        _645_v46 = _dafny.Seq.of(_593_v0);
        let _646_v47;
        _646_v47 = _645_v46;
        _646_v47 = _646_v47;
        let _647_v48;
        let _nw118 = Array((new BigNumber(17)).toNumber());
        _nw118[(_dafny.ZERO).toNumber()] = _593_v0;
        _nw118[(_dafny.ONE).toNumber()] = _593_v0;
        _nw118[(new BigNumber(2)).toNumber()] = new BigNumber(158);
        _nw118[(new BigNumber(3)).toNumber()] = (_593_v0).multipliedBy(_593_v0);
        _nw118[(new BigNumber(4)).toNumber()] = (_645_v46)[_module.__default.safeIndex(_593_v0, new BigNumber((_645_v46).length))];
        _nw118[(new BigNumber(5)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(6)).toNumber()] = new BigNumber(262);
        _nw118[(new BigNumber(7)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(8)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(9)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(10)).toNumber()] = new BigNumber(-672);
        _nw118[(new BigNumber(11)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(12)).toNumber()] = (new BigNumber((_645_v46).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-606)), function (_648_i6) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })).length));
        _nw118[(new BigNumber(13)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(14)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(15)).toNumber()] = _593_v0;
        _nw118[(new BigNumber(16)).toNumber()] = _593_v0;
        _647_v48 = _nw118;
        let _index86 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_647_v48).length));
        (_647_v48)[_index86] = (_dafny.ZERO).minus(_593_v0);
        let _index87 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_647_v48).length));
        let _rhs97 = _593_v0;
        let _rhs98 = (_dafny.ZERO).minus(_593_v0);
        let _lhs72 = _647_v48;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_647_v48).length));
        let _lhs74 = globalState;
        _lhs72[_lhs73] = _rhs97;
        _lhs74.f5 = _rhs98;
        let _649_v49;
        _649_v49 = _dafny.Map.Empty.slice().updateUnsafe(_593_v0,!(!((_this).f13)));
        let _650_v50;
        _650_v50 = _dafny.Seq.of((_this).f13);
        let _651_v51;
        _651_v51 = _module.D0.create_DC0(_649_v49);
        let _652_v52;
        _652_v52 = _dafny.MultiSet.fromElements((_this).f13);
        let _pat_let_tv15 = _647_v48;
        let _pat_let_tv16 = _647_v48;
        let _653_v53;
        let _nw119 = Array((new BigNumber(22)).toNumber());
        _nw119[(_dafny.ZERO).toNumber()] = (_this).f13;
        _nw119[(_dafny.ONE).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((_this).f13), _dafny.Seq.of(_module.__default.fm1((_this).f13, globalState)));
        _nw119[(new BigNumber(2)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(3)).toNumber()] = (((_649_v49).contains(_593_v0)) ? ((_649_v49).get(_593_v0)) : (!((_this).f13)));
        _nw119[(new BigNumber(4)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(5)).toNumber()] = (_593_v0).isLessThan((_647_v48)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_647_v48).length))]);
        _nw119[(new BigNumber(6)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(7)).toNumber()] = !((_this).f13);
        _nw119[(new BigNumber(8)).toNumber()] = !((_this).f13);
        _nw119[(new BigNumber(9)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(10)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(11)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(12)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(13)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(14)).toNumber()] = !((((_649_v49).contains(_593_v0)) ? ((_649_v49).get(_593_v0)) : ((_this).f13)));
        _nw119[(new BigNumber(15)).toNumber()] = !_dafny.areEqual(_650_v50, _dafny.Seq.of((_this).f13, (_this).f13));
        _nw119[(new BigNumber(16)).toNumber()] = false;
        _nw119[(new BigNumber(17)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(18)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(19)).toNumber()] = (_this).f13;
        _nw119[(new BigNumber(20)).toNumber()] = !((_this).f13);
        _nw119[(new BigNumber(21)).toNumber()] = (((_630_v32).fm4(function (_pat_let8_0) {
          return function (_654_dt__update__tmp_h0) {
            return function (_pat_let9_0) {
              return function (_655_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_655_dt__update_hcf0_h0);
              }(_pat_let9_0);
            }(_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv16)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_pat_let_tv15).length))],(_this).f13));
          }(_pat_let8_0);
        }(_651_v51), _651_v51, _652_v52, globalState)) ? (true) : ((_this).f13));
        _653_v53 = _nw119;
        let _index88 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_653_v53).length));
        (_653_v53)[_index88] = (_this).f13;
        let _index89 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_653_v53).length));
        (_653_v53)[_index89] = !((_this).f13);
        let _656_v54;
        _656_v54 = new _dafny.CodePoint('p'.codePointAt(0));
        _656_v54 = _656_v54;
        let _657_v55;
        let _nw120 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
        _657_v55 = _nw120;
        _657_v55 = _657_v55;
      } else {
        let _658_v56;
        let _nw121 = new _module.C1();
        _nw121.__ctor();
        _658_v56 = _nw121;
        (globalState).f2 = (_this).f13;
        let _659_v57;
        _659_v57 = _dafny.Map.Empty.slice().updateUnsafe(_593_v0,new BigNumber(-311));
        let _660_v58;
        _660_v58 = _dafny.Set.fromElements(new BigNumber((_659_v57).length));
        if (((_module.__default.fm30(globalState)).Intersect(_660_v58)).IsProperSubsetOf(_660_v58)) {
          let _661_v59;
          let _init21 = function (_662_i7) {
            return false;
          };
          let _nw122 = Array((new BigNumber(13)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw122.length); _i0_21++) {
            _nw122[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _661_v59 = _nw122;
          _661_v59 = _661_v59;
          let _663_v60;
          let _nw123 = new _module.C1();
          _nw123.__ctor();
          _663_v60 = _nw123;
          (globalState).f0 = (_this).f13;
          let _664_v61;
          let _nw124 = new _module.C0();
          _nw124.__ctor((_this).f13, _661_v59);
          _664_v61 = _nw124;
          _664_v61 = _664_v61;
          let _665_v62;
          let _nw125 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _665_v62 = _nw125;
          let _666_v63;
          _666_v63 = _dafny.MultiSet.fromElements(_665_v62, _665_v62);
          let _667_v64;
          _667_v64 = _dafny.Map.Empty.slice().updateUnsafe(_630_v32,_666_v63);
          let _668_v65;
          _668_v65 = _dafny.Seq.of(_593_v0);
          let _669_v66;
          _669_v66 = _665_v62;
          let _670_v67;
          _670_v67 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_664_v61).f10),new _dafny.CodePoint('v'.codePointAt(0)));
          let _671_v68;
          _671_v68 = _dafny.Seq.of(_666_v63, _666_v63, _666_v63, ((_666_v63).update(_665_v62, _module.__default.abs(new BigNumber((_670_v67).length)))).update(_665_v62, _module.__default.abs(_593_v0)), _666_v63);
          let _672_v69;
          _672_v69 = _module.D2.create_DC6((_this).f13);
          let _673_v70;
          _673_v70 = _dafny.Map.Empty.slice().updateUnsafe((_664_v61).f10,_593_v0);
          let _674_v71;
          _674_v71 = new _dafny.CodePoint('q'.codePointAt(0));
          let _675_v72;
          _675_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm18(_672_v69, new BigNumber((_660_v58).length), new BigNumber((_673_v70).length), globalState),_674_v71);
          let _676_v73;
          let _nw126 = Array((new BigNumber(19)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = _666_v63;
          _nw126[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_665_v62);
          _nw126[(new BigNumber(2)).toNumber()] = ((((_667_v64).contains(_658_v56)) ? ((_667_v64).get(_658_v56)) : (_dafny.MultiSet.fromElements(_665_v62)))).update(_665_v62, _module.__default.abs(new BigNumber((_668_v65).length)));
          _nw126[(new BigNumber(3)).toNumber()] = (_666_v63).update((_669_v66), _module.__default.abs(new BigNumber(89)));
          _nw126[(new BigNumber(4)).toNumber()] = (_666_v63).Union(_666_v63);
          _nw126[(new BigNumber(5)).toNumber()] = (_666_v63).Difference(_666_v63);
          _nw126[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(_665_v62, _665_v62);
          _nw126[(new BigNumber(7)).toNumber()] = _666_v63;
          _nw126[(new BigNumber(8)).toNumber()] = (_671_v68)[_module.__default.safeIndex(_593_v0, new BigNumber((_671_v68).length))];
          _nw126[(new BigNumber(9)).toNumber()] = (_666_v63).Difference((_666_v63).update(_665_v62, _module.__default.abs(new BigNumber((_675_v72).length))));
          _nw126[(new BigNumber(10)).toNumber()] = ((_module.__default.fm1((_this).f13, globalState)) ? (_666_v63) : (_666_v63));
          _nw126[(new BigNumber(11)).toNumber()] = _666_v63;
          _nw126[(new BigNumber(12)).toNumber()] = _666_v63;
          _nw126[(new BigNumber(13)).toNumber()] = _666_v63;
          _nw126[(new BigNumber(14)).toNumber()] = _666_v63;
          _nw126[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements(_665_v62, _665_v62, _665_v62);
          _nw126[(new BigNumber(16)).toNumber()] = (_666_v63).Union(_666_v63);
          _nw126[(new BigNumber(17)).toNumber()] = (_666_v63).Intersect(_dafny.MultiSet.fromElements(_665_v62));
          _nw126[(new BigNumber(18)).toNumber()] = (_671_v68)[_module.__default.safeIndex(_593_v0, new BigNumber((_671_v68).length))];
          _676_v73 = _nw126;
          let _index90 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_676_v73).length));
          (_676_v73)[_index90] = _666_v63;
          let _index91 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_676_v73).length));
          (_676_v73)[_index91] = _666_v63;
        } else {
          let _677_v74;
          _677_v74 = _dafny.Set.fromElements(_659_v57, _659_v57);
          let _678_v75;
          let _nw127 = Array((new BigNumber(5)).toNumber());
          _nw127[(_dafny.ZERO).toNumber()] = _593_v0;
          _nw127[(_dafny.ONE).toNumber()] = (((_this).f13) ? (_593_v0) : (_593_v0));
          _nw127[(new BigNumber(2)).toNumber()] = _593_v0;
          _nw127[(new BigNumber(3)).toNumber()] = _593_v0;
          _nw127[(new BigNumber(4)).toNumber()] = new BigNumber((_677_v74).length);
          _678_v75 = _nw127;
          let _index92 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_678_v75).length));
          (_678_v75)[_index92] = _593_v0;
          let _679_v76;
          _679_v76 = _dafny.Seq.of((_this).f13);
          let _680_v77;
          _680_v77 = _dafny.Seq.of(_679_v76);
          let _index93 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_678_v75).length));
          let _rhs99 = (_dafny.ZERO).minus(_593_v0);
          let _rhs100 = _678_v75;
          let _rhs101 = _module.__default.fm30(globalState);
          let _rhs102 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_680_v77, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-310)), function (_681_i8) {
            return _dafny.Seq.of((_this).f13);
          }))).length));
          let _lhs75 = globalState;
          let _lhs76 = _678_v75;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_678_v75).length));
          _lhs75.f5 = _rhs99;
          _678_v75 = _rhs100;
          _660_v58 = _rhs101;
          _lhs76[_lhs77] = _rhs102;
          let _682_v78;
          _682_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
          let _683_v79;
          _683_v79 = _module.D2.create_DC7(_593_v0, _dafny.Map.Empty.slice().updateUnsafe(_682_v78,_593_v0), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("fjam")).length)));
          let _684_v80;
          _684_v80 = _dafny.MultiSet.fromElements(((!((_this).f13)) ? (_683_v79) : (_683_v79)));
          _684_v80 = (_684_v80).Union(_684_v80);
          let _685_v81;
          _685_v81 = new _dafny.CodePoint('p'.codePointAt(0));
          _685_v81 = _685_v81;
          let _686_v82;
          let _nw128 = new _module.C1();
          _nw128.__ctor();
          _686_v82 = _nw128;
          (globalState).f0 = (_this).f13;
        }
        let _687_v83;
        let _nw129 = Array((new BigNumber(24)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = (_this).f13;
        _nw129[(_dafny.ONE).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(2)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(3)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(4)).toNumber()] = _module.__default.fm1((_this).f13, globalState);
        _nw129[(new BigNumber(5)).toNumber()] = false;
        _nw129[(new BigNumber(6)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(7)).toNumber()] = !((_this).f13);
        _nw129[(new BigNumber(8)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(9)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(10)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(11)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(12)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(13)).toNumber()] = !((_this).f13);
        _nw129[(new BigNumber(14)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(15)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(16)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(17)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(18)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(19)).toNumber()] = !((_this).f13);
        _nw129[(new BigNumber(20)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(21)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(22)).toNumber()] = (_this).f13;
        _nw129[(new BigNumber(23)).toNumber()] = (_this).f13;
        _687_v83 = _nw129;
        let _688_v84;
        let _nw130 = new _module.C0();
        _nw130.__ctor((_this).f13, _687_v83);
        _688_v84 = _nw130;
        let _689_v85;
        _689_v85 = _dafny.Seq.UnicodeFromString("pswtnsaun");
        (globalState).f0 = !(!(false) || (_dafny.Seq.IsProperPrefixOf(_689_v85, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-527)), function (_690_i9) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }))));
      }
      let _691_v86;
      let _init22 = function (_692_i11) {
        return false;
      };
      let _nw131 = Array((new BigNumber(11)).toNumber());
      for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw131.length); _i0_22++) {
        _nw131[_i0_22] = _init22(new BigNumber(_i0_22));
      }
      _691_v86 = _nw131;
      let _693_v87;
      _693_v87 = _dafny.Seq.UnicodeFromString("gvhxsqh");
      let _694_v88;
      _694_v88 = _dafny.Map.Empty.slice().updateUnsafe(_691_v86,_693_v87);
      let _695_v89;
      _695_v89 = _dafny.Seq.of(new BigNumber((_694_v88).length));
      r0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(691)), ((_696_v0) => function (_697_i10) {
        return _696_v0;
      })(_593_v0)), _695_v89);
      let _698_v90;
      _698_v90 = _dafny.MultiSet.fromElements(new BigNumber(507));
      let _699_v91;
      _699_v91 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_698_v90);
      r1 = (((_699_v91).contains(!(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)).isEqualTo(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_593_v0,!((_this).f13))).update(_593_v0, (_this).f13)).length)))) ? ((_699_v91).get(!(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)).isEqualTo(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_593_v0,!((_this).f13))).update(_593_v0, (_this).f13)).length)))) : (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), ((_700_v0) => function (_701_i12) {
        return _700_v0;
      })(_593_v0)))));
      let _702_v92;
      _702_v92 = _dafny.Map.Empty.slice().updateUnsafe(_693_v87,(_this).f13);
      r2 = _702_v92;
      return [r0, r1, r2];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _703_v0;
      let _init23 = function (_704_i0) {
        return (_this).f13;
      };
      let _nw132 = Array((new BigNumber(10)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw132.length); _i0_23++) {
        _nw132[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _703_v0 = _nw132;
      let _705_v1;
      let _nw133 = new _module.C0();
      _nw133.__ctor((_this).f13, _703_v0);
      _705_v1 = _nw133;
      let _706_v2;
      _706_v2 = _module.D2.create_DC6((_705_v1).f10);
      let _707_v3;
      _707_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm18(_706_v2, p1, p1, globalState),!((_this).f13));
      let _708_v4;
      _708_v4 = _dafny.Set.fromElements((_this).f13, (_705_v1).f10);
      _707_v3 = (_707_v3).update((new BigNumber((_708_v4).length)).minus(p1), !((_this).f13));
      let _709_v5;
      _709_v5 = new _dafny.CodePoint('o'.codePointAt(0));
      _709_v5 = _module.__default.fm20(_module.__default.safeDivisionInt(p1, p1), (_this).f13, p1, globalState);
      let _710_i1;
      _710_i1 = _dafny.ZERO;
      L4: {
        while (false) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_710_i1)) {
              break L4;
            }
            _710_i1 = (_710_i1).plus(_dafny.ONE);
            let _711_v6;
            let _init24 = function (_712_i2) {
              return _module.__default.safeModuloInt(_712_i2, new BigNumber(349));
            };
            let _nw134 = Array((new BigNumber(15)).toNumber());
            for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw134.length); _i0_24++) {
              _nw134[_i0_24] = _init24(new BigNumber(_i0_24));
            }
            _711_v6 = _nw134;
            let _713_v7;
            _713_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_711_v6);
            let _714_v8;
            _714_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,_711_v6);
            let _715_v9;
            _715_v9 = _711_v6;
            let _716_v10;
            let _nw135 = Array((new BigNumber(18)).toNumber());
            _nw135[(_dafny.ZERO).toNumber()] = _711_v6;
            _nw135[(_dafny.ONE).toNumber()] = _711_v6;
            _nw135[(new BigNumber(2)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(3)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(4)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(5)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(6)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(7)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(8)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(9)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(10)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(11)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(12)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(13)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(14)).toNumber()] = (((_713_v7).contains((_705_v1).f10)) ? ((_713_v7).get((_705_v1).f10)) : ((((_714_v8).contains(p1)) ? ((_714_v8).get(p1)) : (_711_v6))));
            _nw135[(new BigNumber(15)).toNumber()] = _711_v6;
            _nw135[(new BigNumber(16)).toNumber()] = (_715_v9);
            _nw135[(new BigNumber(17)).toNumber()] = _711_v6;
            _716_v10 = _nw135;
            let _rhs103 = _716_v10;
            let _rhs104 = _703_v0;
            _716_v10 = _rhs103;
            _703_v0 = _rhs104;
            let _717_v11;
            _717_v11 = _dafny.Seq.of(_709_v5, new _dafny.CodePoint('y'.codePointAt(0)), _709_v5);
            _709_v5 = (_717_v11)[_module.__default.safeIndex(new BigNumber(940), new BigNumber((_717_v11).length))];
            let _718_v12;
            _718_v12 = _dafny.Seq.of((_705_v1).f10);
            (globalState).f0 = (((_this).f13) ? ((_718_v12)[_module.__default.safeIndex(p1, new BigNumber((_718_v12).length))]) : (!(_dafny.MultiSet.fromElements((_705_v1).f10, (_705_v1).f10)).contains((_705_v1).f10)));
            (globalState).f0 = (_this).f13;
          }
        }
      }
      let _719_v13;
      _719_v13 = _dafny.Seq.of((_dafny.ZERO).minus(p1), p1, p1);
      let _720_v14;
      _720_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(364),(_dafny.ZERO).minus(p1));
      let _721_v15;
      _721_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_720_v14).length),p1);
      let _722_i3;
      _722_i3 = _dafny.ZERO;
      L5: {
        while (!(_dafny.Seq.IsPrefixOf(_719_v13, _dafny.Seq.update(_dafny.Seq.of(new BigNumber(599)), _module.__default.safeIndex((((_721_v15).contains(new BigNumber(820))) ? ((_721_v15).get(new BigNumber(820))) : ((_719_v13)[_module.__default.safeIndex(p1, new BigNumber((_719_v13).length))])), new BigNumber((_dafny.Seq.of(new BigNumber(599))).length)), p1)))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_722_i3)) {
              break L5;
            }
            _722_i3 = (_722_i3).plus(_dafny.ONE);
            let _723_v16;
            let _nw136 = new _module.C1();
            _nw136.__ctor();
            _723_v16 = _nw136;
            let _724_v17;
            _724_v17 = _dafny.MultiSet.fromElements((_705_v1).f10);
            _724_v17 = (_724_v17).Union(_724_v17);
            (globalState).f5 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, new BigNumber(538)));
            let _725_v18;
            _725_v18 = _dafny.Map.Empty.slice().updateUnsafe(!(false) || ((_this).f13),(p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))]);
            _725_v18 = _725_v18;
          }
        }
      }
      _721_v15 = (_721_v15).update(p1, p1);
      let _726_v19;
      _726_v19 = _dafny.Map.Empty.slice().updateUnsafe(!((_705_v1).f10),new BigNumber(381));
      r0 = _module.__default.safeDivisionInt((p1).minus(p1), (((_726_v19).contains((_this).f13)) ? ((_726_v19).get((_this).f13)) : (p1)));
      r1 = p1;
      let _727_v20;
      _727_v20 = _module.D12.create_DC36((_705_v1).f10, p1, false);
      let _pat_let_tv17 = _705_v1;
      let _pat_let_tv18 = _719_v13;
      let _pat_let_tv19 = _719_v13;
      let _pat_let_tv20 = _705_v1;
      let _pat_let_tv21 = _705_v1;
      let _pat_let_tv22 = _705_v1;
      r2 = function (_source10) {
        if (_source10.is_DC34) {
          let _728___mcc_h0 = (_source10).cf48;
          let _729___mcc_h1 = (_source10).cf49;
          let _730___mcc_h2 = (_source10).cf50;
          let _731_cf50 = _730___mcc_h2;
          let _732_cf49 = _729___mcc_h1;
          let _733_cf48 = _728___mcc_h0;
          return (_pat_let_tv17).f10;
        } else if (_source10.is_DC35) {
          return _dafny.Seq.IsPrefixOf(_pat_let_tv18, _pat_let_tv19);
        } else if (_source10.is_DC36) {
          let _734___mcc_h3 = (_source10).cf51;
          let _735___mcc_h4 = (_source10).cf52;
          let _736___mcc_h5 = (_source10).cf53;
          let _737_cf53 = _736___mcc_h5;
          let _738_cf52 = _735___mcc_h4;
          let _739_cf51 = _734___mcc_h3;
          return !(_738_cf52).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(_module.D0.create_DC1((_pat_let_tv20).f10), _module.D0.create_DC1(_739_cf51), _module.D0.create_DC1(_739_cf51))).cardinality()));
        } else if (_source10.is_DC33) {
          let _740___mcc_h6 = (_source10).cf47;
          let _741_cf47 = _740___mcc_h6;
          return (_pat_let_tv21).f10;
        } else {
          let _742___mcc_h7 = (_source10).cf54;
          let _743_cf54 = _742___mcc_h7;
          return (_pat_let_tv22).f10;
        }
      }(_727_v20);
      return [r0, r1, r2];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor(f12) {
      let _this = this;
      (_this)._f12 = f12;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f12).isLessThanOrEqualTo((_this).f12);
    };
    fm11(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("vd")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f12))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(!(!(false)))))),new BigNumber(-214))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f12)));
    };
    fm12(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_dafny.Seq.of((_this).f12), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f12)).length), (_this).f12)), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of((_this).f12, (_this).f12, new BigNumber((_dafny.Seq.UnicodeFromString("cvi")).length), (_dafny.ZERO).minus((_this).f12)), _dafny.Seq.of((_this).f12), _dafny.Seq.of((_this).f12), _dafny.Seq.of((_this).f12, new BigNumber((_dafny.Seq.of(false)).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-877)), function (_744_i0) {
        return _dafny.Seq.of((_this).f12, (_this).f12);
      })));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _745_v0;
      _745_v0 = true;
      if (!(_745_v0) || (_745_v0)) {
        let _746_v1;
        let _init25 = ((_747_v0) => function (_748_i0) {
          return (_module.D0.create_DC1(_747_v0)).dtor_cf1;
        })(_745_v0);
        let _nw137 = Array((_dafny.ONE).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw137.length); _i0_25++) {
          _nw137[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _746_v1 = _nw137;
        _746_v1 = _746_v1;
        if (_745_v0) {
          let _749_v2;
          _749_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,_745_v0);
          let _750_v3;
          let _nw138 = Array((new BigNumber(6)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = p0;
          _nw138[(_dafny.ONE).toNumber()] = (_this).f12;
          _nw138[(new BigNumber(2)).toNumber()] = new BigNumber((_749_v2).length);
          _nw138[(new BigNumber(3)).toNumber()] = p2;
          _nw138[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("fqcgspfpy")).length));
          _nw138[(new BigNumber(5)).toNumber()] = (_this).f12;
          _750_v3 = _nw138;
          let _751_v4;
          _751_v4 = _750_v3;
          let _752_v5;
          _752_v5 = _dafny.Seq.of(_750_v3, (_750_v3));
          let _753_v6;
          _753_v6 = _dafny.Set.fromElements(_745_v0, _745_v0);
          let _754_v7;
          _754_v7 = _dafny.Set.fromElements(_750_v3, _750_v3, (_752_v5)[_module.__default.safeIndex(new BigNumber((_753_v6).length), new BigNumber((_752_v5).length))]);
          _754_v7 = _754_v7;
          let _755_v8;
          _755_v8 = _dafny.MultiSet.fromElements(_module.__default.fm1(_745_v0, globalState), _745_v0, _745_v0);
          let _756_v9;
          _756_v9 = _dafny.Seq.UnicodeFromString("axeend");
          let _rhs105 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_756_v9, _756_v9), _dafny.Seq.Concat(_756_v9, _756_v9));
          let _rhs106 = _755_v8;
          r0 = _rhs105;
          _755_v8 = _rhs106;
          let _757_v10;
          _757_v10 = new _dafny.CodePoint('u'.codePointAt(0));
          _757_v10 = _757_v10;
          let _index94 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((p1).length));
          (p1)[_index94] = _750_v3;
          let _index95 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((p1).length));
          (p1)[_index95] = _750_v3;
          let _758_v11;
          _758_v11 = _module.D1.create_DC4(p3);
          let _759_v12;
          _759_v12 = _dafny.Map.Empty.slice().updateUnsafe(_758_v11,p3);
          let _760_v13;
          _760_v13 = _dafny.Map.Empty.slice().updateUnsafe(_745_v0,_759_v12);
          _760_v13 = (_dafny.Map.Empty.slice().updateUnsafe(_745_v0,_dafny.Map.Empty.slice().updateUnsafe(_758_v11,new BigNumber(-438)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_745_v0,_759_v12));
        } else {
          let _761_v14;
          _761_v14 = _dafny.Seq.of(_745_v0, !(!(_745_v0)));
          let _762_v15;
          let _nw139 = new _module.C0();
          _nw139.__ctor((_761_v14)[_module.__default.safeIndex(p2, new BigNumber((_761_v14).length))], _746_v1);
          _762_v15 = _nw139;
          let _763_v16;
          let _init26 = ((_764_v0) => function (_765_i1) {
            return (_765_i1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_764_v0,true)).length));
          })(_745_v0);
          let _nw140 = Array((new BigNumber(22)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw140.length); _i0_26++) {
            _nw140[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _763_v16 = _nw140;
          let _index96 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_763_v16).length));
          (_763_v16)[_index96] = p0;
          let _index97 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_763_v16).length));
          (_763_v16)[_index97] = (_this).f12;
          let _766_v17;
          _766_v17 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_745_v0,(_762_v15).f10));
          _766_v17 = _module.__default.fm13(globalState);
          let _767_v18;
          _767_v18 = _module.D4.create_DC14();
          let _rhs107 = (((_762_v15).f10) ? ((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, p2))) : (_module.__default.safeDivisionInt((_this).f12, p2)));
          let _rhs108 = _767_v18;
          let _lhs78 = globalState;
          _lhs78.f5 = _rhs107;
          _767_v18 = _rhs108;
          let _768_v19;
          _768_v19 = _dafny.Seq.UnicodeFromString("uqxmp");
          _768_v19 = ((_745_v0) ? (_dafny.Seq.Concat(_768_v19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-716)), function (_769_i2) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }))) : (_dafny.Seq.Concat(_module.__default.fm14(_768_v19, p3, globalState), _dafny.Seq.UnicodeFromString("fraq"))));
        }
        (globalState).f5 = (_dafny.ZERO).minus((p0).multipliedBy(new BigNumber((_dafny.Seq.of(_745_v0, _745_v0, _745_v0)).length)));
        r1 = p2;
        let _770_v20;
        _770_v20 = _module.D4.create_DC11(_746_v1);
        let _rhs109 = _770_v20;
        let _rhs110 = false;
        let _lhs79 = globalState;
        _770_v20 = _rhs109;
        _lhs79.f0 = _rhs110;
      } else {
        let _771_v21;
        _771_v21 = _dafny.Map.Empty.slice().updateUnsafe(_745_v0,_745_v0);
        let _772_v22;
        _772_v22 = _dafny.Set.fromElements(_745_v0, _745_v0);
        _771_v21 = (_771_v21).update((_772_v22).IsDisjointFrom(_772_v22), _745_v0);
        let _773_v23;
        _773_v23 = _dafny.MultiSet.fromElements(p3, (_this).f12);
        let _774_v24;
        _774_v24 = _dafny.Seq.of(p0);
        let _775_v25;
        _775_v25 = _dafny.Seq.of(_745_v0, _745_v0, _745_v0, _745_v0);
        let _776_v26;
        _776_v26 = _dafny.Map.Empty.slice().updateUnsafe(p3,_745_v0);
        let _777_v27;
        _777_v27 = _module.D0.create_DC0(_776_v26);
        let _778_v28;
        _778_v28 = _dafny.MultiSet.fromElements(false);
        let _779_v29;
        _779_v29 = _dafny.Seq.of(_745_v0, (_775_v25)[_module.__default.safeIndex(p2, new BigNumber((_775_v25).length))], (_this).fm4(_777_v27, _module.D0.create_DC0(_776_v26), _778_v28, globalState), true);
        let _rhs111 = (_dafny.MultiSet.FromArray(_774_v24)).IsSubsetOf(_773_v23);
        let _rhs112 = (_module.__default.safeModuloInt(p3, (_this).f12)).plus(new BigNumber((_dafny.Seq.Concat(_779_v29, _779_v29)).length));
        let _lhs80 = globalState;
        let _lhs81 = globalState;
        _lhs80.f0 = _rhs111;
        _lhs81.f5 = _rhs112;
        let _780_v30;
        _780_v30 = _module.D0.create_DC1(_745_v0);
        let _source11 = _780_v30;
        if (_source11.is_DC1) {
          let _781___mcc_h0 = (_source11).cf1;
          let _782_cf1 = _781___mcc_h0;
          (globalState).f0 = _782_cf1;
          let _783_v31;
          _783_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(52),new BigNumber((_779_v29).length));
          let _784_v32;
          let _nw141 = Array((new BigNumber(19)).toNumber());
          _nw141[(_dafny.ZERO).toNumber()] = !((_780_v30).dtor_cf1);
          _nw141[(_dafny.ONE).toNumber()] = _782_cf1;
          _nw141[(new BigNumber(2)).toNumber()] = _745_v0;
          _nw141[(new BigNumber(3)).toNumber()] = _745_v0;
          _nw141[(new BigNumber(4)).toNumber()] = _782_cf1;
          _nw141[(new BigNumber(5)).toNumber()] = _745_v0;
          _nw141[(new BigNumber(6)).toNumber()] = true;
          _nw141[(new BigNumber(7)).toNumber()] = (_this).fm4(_777_v27, _777_v27, _778_v28, globalState);
          _nw141[(new BigNumber(8)).toNumber()] = _782_cf1;
          _nw141[(new BigNumber(9)).toNumber()] = _745_v0;
          _nw141[(new BigNumber(10)).toNumber()] = _745_v0;
          _nw141[(new BigNumber(11)).toNumber()] = _782_cf1;
          _nw141[(new BigNumber(12)).toNumber()] = _782_cf1;
          _nw141[(new BigNumber(13)).toNumber()] = _745_v0;
          _nw141[(new BigNumber(14)).toNumber()] = _782_cf1;
          _nw141[(new BigNumber(15)).toNumber()] = _745_v0;
          _nw141[(new BigNumber(16)).toNumber()] = _782_cf1;
          _nw141[(new BigNumber(17)).toNumber()] = (_this).fm4(_777_v27, _777_v27, _778_v28, globalState);
          _nw141[(new BigNumber(18)).toNumber()] = _782_cf1;
          _784_v32 = _nw141;
          let _785_v33;
          let _nw142 = new _module.C0();
          _nw142.__ctor(((((_783_v31).contains(p2)) ? ((_783_v31).get(p2)) : (new BigNumber(-979)))).isLessThan(p2), _784_v32);
          _785_v33 = _nw142;
          let _786_v34;
          _786_v34 = _dafny.Seq.UnicodeFromString("gvwujy");
          _786_v34 = _786_v34;
          let _787_v35;
          let _nw143 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _787_v35 = _nw143;
          _787_v35 = _787_v35;
        } else if (_source11.is_DC0) {
          let _788___mcc_h1 = (_source11).cf0;
          let _789_cf0 = _788___mcc_h1;
          let _790_v36;
          let _nw144 = Array((new BigNumber(29)).toNumber());
          _nw144[(_dafny.ZERO).toNumber()] = _745_v0;
          _nw144[(_dafny.ONE).toNumber()] = _745_v0;
          _nw144[(new BigNumber(2)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(3)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(4)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(5)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(6)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(7)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(8)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(9)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(10)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(11)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(12)).toNumber()] = false;
          _nw144[(new BigNumber(13)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(14)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(15)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(16)).toNumber()] = true;
          _nw144[(new BigNumber(17)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(18)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(19)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(20)).toNumber()] = !(_745_v0);
          _nw144[(new BigNumber(21)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(22)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(23)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(24)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(25)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(26)).toNumber()] = (_775_v25)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_775_v25).length))];
          _nw144[(new BigNumber(27)).toNumber()] = _745_v0;
          _nw144[(new BigNumber(28)).toNumber()] = _745_v0;
          _790_v36 = _nw144;
          let _791_v37;
          _791_v37 = _dafny.Map.Empty.slice().updateUnsafe(_790_v36,_745_v0);
          _791_v37 = (_791_v37).Merge(_791_v37);
          let _792_v38;
          _792_v38 = new _dafny.CodePoint('x'.codePointAt(0));
          let _793_v40;
          _793_v40 = _dafny.Map.Empty.slice().updateUnsafe(_745_v0,(_this).f12);
          let _794_v41;
          _794_v41 = _dafny.Seq.of(_793_v40);
          _module.__default.m0(_745_v0, ((_745_v0) ? (_792_v38) : (_792_v38)), new BigNumber((function () {
            let _coll39 = new _dafny.Set();
            for (const _compr_39 of _dafny.IntegerRange(new BigNumber(364), new BigNumber(751))) {
              let _795_v39 = _compr_39;
              if (((new BigNumber(364)).isLessThanOrEqualTo(_795_v39)) && ((_795_v39).isLessThan(new BigNumber(751)))) {
                _coll39.add((_795_v39).plus(new BigNumber(165)));
              }
            }
            return _coll39;
          }()).length), (_793_v40).Merge((_794_v41)[_module.__default.safeIndex((_this).f12, new BigNumber((_794_v41).length))]), globalState);
          let _796_v42;
          let _nw145 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _796_v42 = _nw145;
          let _797_v43;
          _797_v43 = _module.D1.create_DC3(_796_v42);
          let _pat_let_tv23 = _796_v42;
          _797_v43 = function (_pat_let10_0) {
            return function (_798_dt__update__tmp_h2) {
              return function (_pat_let11_0) {
                return function (_799_dt__update_hcf3_h0) {
                  return _module.D1.create_DC3(_799_dt__update_hcf3_h0);
                }(_pat_let11_0);
              }(_pat_let_tv23);
            }(_pat_let10_0);
          }(_797_v43);
          (globalState).f0 = ((_745_v0) ? (_745_v0) : (_745_v0));
        } else {
          let _800___mcc_h2 = (_source11).cf2;
          let _801_cf2 = _800___mcc_h2;
          let _802_v44;
          let _nw146 = Array((new BigNumber(23)).toNumber()).fill(_dafny.MultiSet.Empty);
          _802_v44 = _nw146;
          let _803_v45;
          _803_v45 = _dafny.MultiSet.fromElements(_779_v29, _779_v29, _775_v25);
          let _index98 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_802_v44).length));
          (_802_v44)[_index98] = (_803_v45).Intersect((_dafny.MultiSet.FromArray(_dafny.Seq.of(_779_v29, _dafny.Seq.of(_745_v0)))).update(_775_v25, _module.__default.abs(p0)));
          let _804_v46;
          _804_v46 = _dafny.Seq.UnicodeFromString("ofcufvnr");
          let _805_v47;
          _805_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_804_v46).length),p0);
          let _806_v48;
          _806_v48 = _dafny.Seq.of(_805_v47);
          let _807_v49;
          let _nw147 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _807_v49 = _nw147;
          let _index99 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length));
          (p1)[_index99] = _807_v49;
          let _808_v50;
          _808_v50 = _dafny.Map.Empty.slice().updateUnsafe(_745_v0,p2);
          let _index100 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_802_v44).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length));
          let _rhs113 = new BigNumber((_808_v50).length);
          let _rhs114 = _803_v45;
          let _rhs115 = _dafny.Seq.of(_805_v47);
          let _rhs116 = _807_v49;
          let _rhs117 = (_this).f12;
          let _lhs82 = _802_v44;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_802_v44).length));
          let _lhs84 = p1;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length));
          let _lhs86 = globalState;
          r1 = _rhs113;
          _lhs82[_lhs83] = _rhs114;
          _806_v48 = _rhs115;
          _lhs84[_lhs85] = _rhs116;
          _lhs86.f5 = _rhs117;
          let _809_v51;
          let _nw148 = Array((new BigNumber(23)).toNumber());
          _nw148[(_dafny.ZERO).toNumber()] = _774_v24;
          _nw148[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_774_v24, _774_v24);
          _nw148[(new BigNumber(2)).toNumber()] = ((_745_v0) ? (_774_v24) : (_module.__default.fm15(_745_v0, _745_v0, globalState)));
          _nw148[(new BigNumber(3)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_774_v24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-462)), ((_810_p0) => function (_811_i3) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), ((_812_p0) => function (_813_i4) {
              return _812_p0;
            })(_810_p0))).length);
          })(p0)));
          _nw148[(new BigNumber(5)).toNumber()] = _dafny.Seq.of((_this).f12);
          _nw148[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(56)), ((_814_p3) => function (_815_i5) {
            return _814_p3;
          })(p3));
          _nw148[(new BigNumber(7)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(679)), ((_816_p0) => function (_817_i6) {
            return _816_p0;
          })(p0)), _774_v24);
          _nw148[(new BigNumber(9)).toNumber()] = _dafny.Seq.of((_this).f12, p3, _module.__default.fm0(new BigNumber(499), new BigNumber((_805_v47).length), p3, _745_v0, globalState));
          _nw148[(new BigNumber(10)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(11)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(12)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(13)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(985)), ((_818_p3) => function (_819_i7) {
            return _818_p3;
          })(p3));
          _nw148[(new BigNumber(15)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(16)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(17)).toNumber()] = _dafny.Seq.of((_this).f12, p2);
          _nw148[(new BigNumber(18)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), ((_820_v50, _821_p2) => function (_822_i8) {
            return (((_820_v50).contains(false)) ? ((_820_v50).get(false)) : (_821_p2));
          })(_808_v50, p2)), _774_v24);
          _nw148[(new BigNumber(20)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(21)).toNumber()] = _774_v24;
          _nw148[(new BigNumber(22)).toNumber()] = _774_v24;
          _809_v51 = _nw148;
          let _index102 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_809_v51).length));
          (_809_v51)[_index102] = _774_v24;
          let _index103 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_809_v51).length));
          (_809_v51)[_index103] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(939)), ((_823_p2) => function (_824_i9) {
            return _823_p2;
          })(p2));
          let _825_v52;
          let _init27 = ((_826_v0) => function (_827_i10) {
            return _826_v0;
          })(_745_v0);
          let _nw149 = Array((new BigNumber(25)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw149.length); _i0_27++) {
            _nw149[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _825_v52 = _nw149;
          _825_v52 = _825_v52;
          let _828_v53;
          _828_v53 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))],_745_v0);
          let _index104 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_807_v49).length));
          (_807_v49)[_index104] = _module.__default.fm0(p2, p3, p0, (((_828_v53).contains(_807_v49)) ? ((_828_v53).get(_807_v49)) : (_745_v0)), globalState);
          let _index105 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_807_v49).length));
          (_807_v49)[_index105] = (_this).f12;
        }
        let _829_v54;
        _829_v54 = _module.D0.create_DC0(_776_v26);
        let _830_v55;
        _830_v55 = _module.D0.create_DC2(_829_v54);
        let _831_v56;
        _831_v56 = _dafny.MultiSet.fromElements(_830_v55);
        let _832_v57;
        _832_v57 = _dafny.Seq.UnicodeFromString("sfcvgsxqa");
        let _833_v58;
        _833_v58 = _dafny.Seq.of(_775_v25, _dafny.Seq.of(_745_v0, _745_v0), _779_v29, _779_v29, _779_v29);
        let _834_v59;
        _834_v59 = _module.D2.create_DC6(_745_v0);
        let _835_v60;
        let _nw150 = Array((new BigNumber(18)).toNumber());
        _nw150[(_dafny.ZERO).toNumber()] = (_831_v56).IsDisjointFrom(_831_v56);
        _nw150[(_dafny.ONE).toNumber()] = _745_v0;
        _nw150[(new BigNumber(2)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(3)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(4)).toNumber()] = _dafny.areEqual(_774_v24, _dafny.Seq.of(new BigNumber((_779_v29).length), p3, new BigNumber((_832_v57).length)));
        _nw150[(new BigNumber(5)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(6)).toNumber()] = !(_745_v0);
        _nw150[(new BigNumber(7)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_833_v58, _dafny.Seq.of(_775_v25, _779_v29));
        _nw150[(new BigNumber(8)).toNumber()] = false;
        _nw150[(new BigNumber(9)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(10)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(11)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(12)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(13)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(14)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(15)).toNumber()] = (_834_v59).dtor_cf6;
        _nw150[(new BigNumber(16)).toNumber()] = _745_v0;
        _nw150[(new BigNumber(17)).toNumber()] = _745_v0;
        _835_v60 = _nw150;
        let _index106 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_835_v60).length));
        (_835_v60)[_index106] = (_745_v0) || (_745_v0);
        let _index107 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_835_v60).length));
        let _rhs118 = _745_v0;
        let _rhs119 = p3;
        let _rhs120 = (_dafny.ZERO).minus((_dafny.ZERO).minus((p2).minus((p0).multipliedBy((_this).f12))));
        let _rhs121 = (((_834_v59).dtor_cf6) ? (p0) : (p2));
        let _lhs87 = _835_v60;
        let _lhs88 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_835_v60).length));
        let _lhs89 = globalState;
        let _lhs90 = globalState;
        let _lhs91 = globalState;
        _lhs87[_lhs88] = _rhs118;
        _lhs89.f5 = _rhs119;
        _lhs90.f5 = _rhs120;
        _lhs91.f5 = _rhs121;
        let _836_v61;
        _836_v61 = _dafny.Map.Empty.slice().updateUnsafe((_835_v60)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_835_v60).length))],p0);
        let _837_v62;
        _837_v62 = _dafny.Set.fromElements(new BigNumber((_771_v21).length));
        let _838_v63;
        _838_v63 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_837_v62).length)),new BigNumber(-557));
        let _839_v64;
        let _nw151 = Array((new BigNumber(13)).toNumber());
        _nw151[(_dafny.ZERO).toNumber()] = ((_this).f12).multipliedBy(p2);
        _nw151[(_dafny.ONE).toNumber()] = (_this).f12;
        _nw151[(new BigNumber(2)).toNumber()] = new BigNumber((_773_v23).cardinality());
        _nw151[(new BigNumber(3)).toNumber()] = p0;
        _nw151[(new BigNumber(4)).toNumber()] = p0;
        _nw151[(new BigNumber(5)).toNumber()] = p2;
        _nw151[(new BigNumber(6)).toNumber()] = (((_836_v61).contains(_745_v0)) ? ((_836_v61).get(_745_v0)) : ((((_838_v63).contains(new BigNumber((_776_v26).length))) ? ((_838_v63).get(new BigNumber((_776_v26).length))) : (p0))));
        _nw151[(new BigNumber(7)).toNumber()] = p0;
        _nw151[(new BigNumber(8)).toNumber()] = new BigNumber(687);
        _nw151[(new BigNumber(9)).toNumber()] = new BigNumber(448);
        _nw151[(new BigNumber(10)).toNumber()] = p3;
        _nw151[(new BigNumber(11)).toNumber()] = new BigNumber(952);
        _nw151[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(p2);
        _839_v64 = _nw151;
        let _rhs122 = _module.__default.safeDivisionInt(p0, p0);
        let _rhs123 = (((_835_v60)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_835_v60).length))]) ? (new BigNumber(807)) : ((p3).plus(new BigNumber((_dafny.MultiSet.fromElements((_835_v60)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_835_v60).length))])).cardinality()))));
        let _rhs124 = _839_v64;
        let _rhs125 = !(_745_v0);
        let _lhs92 = globalState;
        let _lhs93 = globalState;
        _lhs92.f5 = _rhs122;
        _lhs93.f5 = _rhs123;
        _839_v64 = _rhs124;
        _745_v0 = _rhs125;
      }
      let _840_v65;
      _840_v65 = new _dafny.CodePoint('u'.codePointAt(0));
      _840_v65 = _module.__default.fm16(globalState);
      let _841_v66;
      _841_v66 = _dafny.MultiSet.fromElements(true, false);
      r1 = ((_this).f12).minus(new BigNumber((_841_v66).cardinality()));
      let _842_i11;
      _842_i11 = _dafny.ZERO;
      L6: {
        while (true) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_842_i11)) {
              break L6;
            }
            _842_i11 = (_842_i11).plus(_dafny.ONE);
            if (false) {
              let _843_v67;
              _843_v67 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_745_v0);
              let _844_v68;
              _844_v68 = _dafny.Seq.of(false);
              let _845_v69;
              _845_v69 = _dafny.Seq.UnicodeFromString("kdak");
              let _846_v70;
              _846_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p2,_845_v69),false);
              let _847_v72;
              let _nw152 = Array((new BigNumber(28)).toNumber());
              _nw152[(_dafny.ZERO).toNumber()] = true;
              _nw152[(_dafny.ONE).toNumber()] = _745_v0;
              _nw152[(new BigNumber(2)).toNumber()] = false;
              _nw152[(new BigNumber(3)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(4)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(5)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(6)).toNumber()] = (((_843_v67).contains(new BigNumber((_dafny.Seq.update(_844_v68, _module.__default.safeIndex(p2, new BigNumber((_844_v68).length)), _745_v0)).length))) ? ((_843_v67).get(new BigNumber((_dafny.Seq.update(_844_v68, _module.__default.safeIndex(p2, new BigNumber((_844_v68).length)), _745_v0)).length))) : (_745_v0));
              _nw152[(new BigNumber(7)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(8)).toNumber()] = _module.__default.fm1(_745_v0, globalState);
              _nw152[(new BigNumber(9)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(10)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(11)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(12)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(13)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(14)).toNumber()] = true;
              _nw152[(new BigNumber(15)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(16)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(17)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(18)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(19)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(20)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(21)).toNumber()] = !(_745_v0);
              _nw152[(new BigNumber(22)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(23)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(24)).toNumber()] = _745_v0;
              _nw152[(new BigNumber(25)).toNumber()] = _module.__default.fm1(_745_v0, globalState);
              _nw152[(new BigNumber(26)).toNumber()] = (((_846_v70).contains(function () {
                let _coll41 = new _dafny.Map();
                for (const _compr_41 of (_843_v67).Keys.Elements) {
                  let _849_v71 = _compr_41;
                  if ((_843_v67).contains(_849_v71)) {
                    _coll41.push([(_849_v71).multipliedBy(p0),_845_v69]);
                  }
                }
                return _coll41;
              }())) ? ((_846_v70).get(function () {
                let _coll40 = new _dafny.Map();
                for (const _compr_40 of (_843_v67).Keys.Elements) {
                  let _848_v71 = _compr_40;
                  if ((_843_v67).contains(_848_v71)) {
                    _coll40.push([(_848_v71).multipliedBy(p0),_845_v69]);
                  }
                }
                return _coll40;
              }())) : (_745_v0));
              _nw152[(new BigNumber(27)).toNumber()] = _745_v0;
              _847_v72 = _nw152;
              let _850_v73;
              let _nw153 = new _module.C0();
              _nw153.__ctor(true, _847_v72);
              _850_v73 = _nw153;
              let _851_v75;
              _851_v75 = _module.D0.create_DC0(function () {
  let _coll42 = new _dafny.Map();
  for (const _compr_42 of _dafny.IntegerRange(new BigNumber(725), new BigNumber(-252))) {
    let _852_v74 = _compr_42;
    if (((new BigNumber(725)).isLessThanOrEqualTo(_852_v74)) && ((_852_v74).isLessThan(new BigNumber(-252)))) {
      _coll42.push([_module.__default.safeDivisionInt(_852_v74, p3),(_850_v73).f10]);
    }
  }
  return _coll42;
}());
              let _pat_let_tv24 = _843_v67;
              _847_v72 = (((_this).fm4(function (_pat_let12_0) {
                return function (_853_dt__update__tmp_h3) {
                  return function (_pat_let13_0) {
                    return function (_854_dt__update_hcf0_h0) {
                      return _module.D0.create_DC0(_854_dt__update_hcf0_h0);
                    }(_pat_let13_0);
                  }(_pat_let_tv24);
                }(_pat_let12_0);
              }(_851_v75), _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(p2,_745_v0)), _dafny.MultiSet.FromArray(_844_v68), globalState)) ? ((_850_v73).f11) : (_847_v72));
              let _855_v76;
              _855_v76 = _dafny.Set.fromElements((_850_v73).f10);
              (globalState).f5 = (new BigNumber((_855_v76).length)).multipliedBy(p3);
              r0 = ((_855_v76).Difference(_855_v76)).IsSubsetOf(_855_v76);
              let _856_v77;
              let _nw154 = new _module.C0();
              _nw154.__ctor(_745_v0, (_850_v73).f11);
              _856_v77 = _nw154;
            } else {
              let _857_v78;
              let _init28 = ((_858_v0) => function (_859_i12) {
                return _858_v0;
              })(_745_v0);
              let _nw155 = Array((new BigNumber(4)).toNumber());
              for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw155.length); _i0_28++) {
                _nw155[_i0_28] = _init28(new BigNumber(_i0_28));
              }
              _857_v78 = _nw155;
              let _860_v79;
              let _nw156 = new _module.C0();
              _nw156.__ctor(_745_v0, _857_v78);
              _860_v79 = _nw156;
              let _861_v80;
              let _nw157 = Array((new BigNumber(4)).toNumber());
              _nw157[(_dafny.ZERO).toNumber()] = p2;
              _nw157[(_dafny.ONE).toNumber()] = p0;
              _nw157[(new BigNumber(2)).toNumber()] = p0;
              _nw157[(new BigNumber(3)).toNumber()] = p0;
              _861_v80 = _nw157;
              let _862_v81;
              _862_v81 = _dafny.Map.Empty.slice().updateUnsafe(_745_v0,_861_v80);
              let _863_v82;
              _863_v82 = _dafny.Seq.UnicodeFromString("dypeuvr");
              let _864_v83;
              _864_v83 = _dafny.Set.fromElements((_863_v82)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("etmifa")).length), new BigNumber((_863_v82).length))]);
              let _865_v84;
              _865_v84 = _dafny.Map.Empty.slice().updateUnsafe(_745_v0,new BigNumber((_dafny.Seq.update(_863_v82, _module.__default.safeIndex(p3, new BigNumber((_863_v82).length)), _840_v65)).length));
              let _866_v85;
              let _nw158 = Array((new BigNumber(23)).toNumber());
              _nw158[(_dafny.ZERO).toNumber()] = new BigNumber(-539);
              _nw158[(_dafny.ONE).toNumber()] = (new BigNumber((_862_v81).length)).multipliedBy((_dafny.ZERO).minus(p0));
              _nw158[(new BigNumber(2)).toNumber()] = p2;
              _nw158[(new BigNumber(3)).toNumber()] = p0;
              _nw158[(new BigNumber(4)).toNumber()] = p0;
              _nw158[(new BigNumber(5)).toNumber()] = ((_745_v0) ? ((_this).f12) : (p0));
              _nw158[(new BigNumber(6)).toNumber()] = _module.__default.fm0(_module.__default.fm0(p0, p0, new BigNumber((_module.__default.fm17(true, p3, _745_v0, p0, globalState)).length), _745_v0, globalState), (_this).f12, new BigNumber((_864_v83).length), _745_v0, globalState);
              _nw158[(new BigNumber(7)).toNumber()] = ((_745_v0) ? (new BigNumber((_865_v84).length)) : ((_dafny.ZERO).minus((_this).f12)));
              _nw158[(new BigNumber(8)).toNumber()] = (_this).f12;
              _nw158[(new BigNumber(9)).toNumber()] = (_this).f12;
              _nw158[(new BigNumber(10)).toNumber()] = p3;
              _nw158[(new BigNumber(11)).toNumber()] = p0;
              _nw158[(new BigNumber(12)).toNumber()] = p3;
              _nw158[(new BigNumber(13)).toNumber()] = (p2).multipliedBy(p2);
              _nw158[(new BigNumber(14)).toNumber()] = new BigNumber(((((_860_v79).f10) ? (_dafny.Seq.update(_dafny.Seq.of(new BigNumber(568)), _module.__default.safeIndex(new BigNumber(845), new BigNumber((_dafny.Seq.of(new BigNumber(568))).length)), new BigNumber((_863_v82).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), ((_867_p2) => function (_868_i13) {
                return _867_p2;
              })(p2))))).length);
              _nw158[(new BigNumber(15)).toNumber()] = p0;
              _nw158[(new BigNumber(16)).toNumber()] = new BigNumber((_841_v66).cardinality());
              _nw158[(new BigNumber(17)).toNumber()] = p3;
              _nw158[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_863_v82).length), (_this).f12);
              _nw158[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt(p2, p2);
              _nw158[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_this).f12)).cardinality());
              _nw158[(new BigNumber(21)).toNumber()] = _module.__default.safeModuloInt((_this).f12, (_this).f12);
              _nw158[(new BigNumber(22)).toNumber()] = new BigNumber(-487);
              _866_v85 = _nw158;
              _866_v85 = _866_v85;
              (globalState).f5 = (_this).f12;
              (globalState).f0 = !(!(!((_this).fm12(_863_v82, (_this).f12, p0, globalState))) || ((_860_v79).f10));
              (globalState).f0 = _745_v0;
            }
            if (_745_v0) {
              let _869_v86;
              let _nw159 = Array((new BigNumber(11)).toNumber()).fill(false);
              _869_v86 = _nw159;
              _869_v86 = _869_v86;
              let _870_v87;
              _870_v87 = _dafny.Seq.UnicodeFromString("jqxxg");
              let _871_v88;
              let _nw160 = Array((new BigNumber(15)).toNumber());
              _871_v88 = _nw160;
              let _872_v89;
              let _873_v90;
              let _out4;
              let _out5;
              let _outcollector1 = (_this).m3(_module.__default.fm14(_870_v87, p2, globalState), _745_v0, _871_v88, p0, globalState);
              _out4 = _outcollector1[0];
              _out5 = _outcollector1[1];
              _872_v89 = _out4;
              _873_v90 = _out5;
              _870_v87 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("u"), _870_v87), _870_v87);
              let _874_v91;
              _874_v91 = _dafny.Seq.of(new BigNumber((_870_v87).length));
              let _875_v92;
              _875_v92 = _dafny.Map.Empty.slice().updateUnsafe(_874_v91,p2);
              _875_v92 = ((_875_v92).Merge(_875_v92)).Merge(_875_v92);
              let _index108 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_869_v86).length));
              (_869_v86)[_index108] = ((_745_v0) ? (_745_v0) : (true));
              let _876_v93;
              let _nw161 = new _module.C2();
              _nw161.__ctor(false);
              _876_v93 = _nw161;
              let _877_v94;
              _877_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_876_v93);
              let _878_v95;
              _878_v95 = _dafny.MultiSet.fromElements(_872_v89, new BigNumber(-22), new BigNumber(181), new BigNumber((_877_v94).length), new BigNumber(-526));
              let _879_v96;
              _879_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(192)), ((_880_p3) => function (_881_i14) {
                return _880_p3;
              })(p3))).length),_745_v0);
              let _882_v97;
              _882_v97 = _module.D0.create_DC0(_879_v96);
              let _883_v98;
              _883_v98 = _module.D0.create_DC2(_882_v97);
              let _884_v99;
              _884_v99 = _module.D0.create_DC0(_879_v96);
              let _885_v100;
              _885_v100 = _module.D12.create_DC34(_883_v98, _745_v0, (_876_v93).fm4(_884_v99, _884_v99, _841_v66, globalState));
              let _886_v101;
              _886_v101 = _module.D12.create_DC34(_module.D0.create_DC2(_882_v97), _745_v0, !((_885_v100).dtor_cf50));
              let _887_v102;
              _887_v102 = _dafny.Seq.of(_745_v0, (_878_v95).IsDisjointFrom(_dafny.MultiSet.FromArray(_874_v91)), (_885_v100).dtor_cf49, true);
              let _index109 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_869_v86).length));
              (_869_v86)[_index109] = (_887_v102)[_module.__default.safeIndex(_872_v89, new BigNumber((_887_v102).length))];
            } else {
              let _888_v103;
              _888_v103 = _dafny.Map.Empty.slice().updateUnsafe(_745_v0,_745_v0);
              let _889_v104;
              _889_v104 = _dafny.Seq.of(_888_v103);
              let _890_v105;
              _890_v105 = _dafny.Map.Empty.slice().updateUnsafe((_889_v104)[_module.__default.safeIndex(p0, new BigNumber((_889_v104).length))],new BigNumber(762));
              let _891_v106;
              _891_v106 = _module.D2.create_DC7(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("heyc"), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f12), new BigNumber((_dafny.Seq.UnicodeFromString("heyc")).length)), _840_v65)).length), (_890_v105).update(_888_v103, (_this).f12), (_this).f12);
              (globalState).f5 = (_891_v106).dtor_cf9;
              let _892_v107;
              _892_v107 = _dafny.Seq.of(!(_745_v0));
              let _893_v108;
              _893_v108 = _module.D9.create_DC24(true, true, (_892_v107)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f12), new BigNumber((_892_v107).length))]);
              let _894_v109;
              _894_v109 = _dafny.Map.Empty.slice().updateUnsafe(p3,_745_v0);
              let _895_v110;
              _895_v110 = _dafny.Map.Empty.slice().updateUnsafe((_892_v107)[_module.__default.safeIndex(new BigNumber(534), new BigNumber((_892_v107).length))],p0);
              let _rhs126 = ((((((_894_v109).contains(new BigNumber(632))) ? ((_894_v109).get(new BigNumber(632))) : ((((_888_v103).contains(false)) ? ((_888_v103).get(false)) : (_745_v0))))) === (_745_v0)) ? (_module.D9.create_DC24(false, _745_v0, _745_v0)) : (_893_v108));
              let _rhs127 = _module.__default.safeDivisionInt(new BigNumber(((_module.D13.create_DC38(_895_v110)).dtor_cf55).length), _module.__default.safeDivisionInt(new BigNumber((_892_v107).length), p3));
              _893_v108 = _rhs126;
              r1 = _rhs127;
              let _896_v111;
              _896_v111 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f12);
              let _897_v112;
              _897_v112 = _dafny.Map.Empty.slice().updateUnsafe(_896_v111,_745_v0);
              (globalState).f0 = !(!(!((((_897_v112).contains(_896_v111)) ? ((_897_v112).get(_896_v111)) : (_745_v0)))));
              let _898_v113;
              _898_v113 = _dafny.Seq.of((_this).f12);
              let _899_v114;
              _899_v114 = _dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_898_v113).length), p0, p0, _745_v0, globalState), p0);
              let _900_v116;
              let _nw162 = Array((new BigNumber(9)).toNumber());
              _nw162[(_dafny.ZERO).toNumber()] = _745_v0;
              _nw162[(_dafny.ONE).toNumber()] = true;
              _nw162[(new BigNumber(2)).toNumber()] = _dafny.areEqual(_898_v113, _898_v113);
              _nw162[(new BigNumber(3)).toNumber()] = _745_v0;
              _nw162[(new BigNumber(4)).toNumber()] = !(_745_v0);
              _nw162[(new BigNumber(5)).toNumber()] = _745_v0;
              _nw162[(new BigNumber(6)).toNumber()] = _745_v0;
              _nw162[(new BigNumber(7)).toNumber()] = _745_v0;
              _nw162[(new BigNumber(8)).toNumber()] = (_899_v114).IsDisjointFrom(function () {
                let _coll43 = new _dafny.Set();
                for (const _compr_43 of _dafny.IntegerRange(new BigNumber(750), new BigNumber(-18))) {
                  let _901_v115 = _compr_43;
                  if (((new BigNumber(750)).isLessThanOrEqualTo(_901_v115)) && ((_901_v115).isLessThan(new BigNumber(-18)))) {
                    _coll43.add(_module.__default.safeModuloInt(_901_v115, p0));
                  }
                }
                return _coll43;
              }());
              _900_v116 = _nw162;
              let _index110 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_900_v116).length));
              (_900_v116)[_index110] = _745_v0;
              let _902_v117;
              let _nw163 = Array((new BigNumber(16)).toNumber()).fill([]);
              _902_v117 = _nw163;
              let _903_v118;
              let _init29 = ((_904_p3) => function (_905_i15) {
                return _module.__default.safeModuloInt(_905_i15, _904_p3);
              })(p3);
              let _nw164 = Array((new BigNumber(11)).toNumber());
              for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw164.length); _i0_29++) {
                _nw164[_i0_29] = _init29(new BigNumber(_i0_29));
              }
              _903_v118 = _nw164;
              let _906_v119;
              _906_v119 = _903_v118;
              let _index111 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_902_v117).length));
              (_902_v117)[_index111] = _906_v119;
              let _index112 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_900_v116).length));
              let _index113 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_902_v117).length));
              let _rhs128 = !((_745_v0) === (!(_745_v0))) || (_745_v0);
              let _rhs129 = false;
              let _rhs130 = _906_v119;
              let _lhs94 = _900_v116;
              let _lhs95 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_900_v116).length));
              let _lhs96 = _902_v117;
              let _lhs97 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_902_v117).length));
              r2 = _rhs128;
              _lhs94[_lhs95] = _rhs129;
              _lhs96[_lhs97] = _rhs130;
              let _907_v120;
              _907_v120 = _module.D3.create_DC9(new BigNumber(409));
              let _908_v121;
              _908_v121 = _module.D3.create_DC10(_907_v120);
              let _909_v122;
              _909_v122 = _dafny.Seq.UnicodeFromString("fxm");
              let _pat_let_tv25 = _909_v122;
              let _910_v123;
              _910_v123 = _dafny.Map.Empty.slice().updateUnsafe(p3,function (_pat_let14_0) {
                return function (_911_dt__update__tmp_h4) {
                  return function (_pat_let15_0) {
                    return function (_912_dt__update_hcf12_h0) {
                      return _module.D3.create_DC10(_912_dt__update_hcf12_h0);
                    }(_pat_let15_0);
                  }(_module.D3.create_DC8(_pat_let_tv25));
                }(_pat_let14_0);
              }(_908_v121));
              (globalState).f0 = (p0).isLessThan(new BigNumber((_910_v123).length));
            }
            let _913_v124;
            _913_v124 = _module.D11.create_DC30();
            let _914_v125;
            _914_v125 = _module.D11.create_DC32(_913_v124);
            let _915_v126;
            _915_v126 = _module.D11.create_DC32((_914_v125).dtor_cf46);
            let _916_v128;
            let _init30 = ((_917_v66, _918_p2) => function (_919_i16) {
              return _module.D4.create_DC12(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f12, new BigNumber((function () {
  let _coll44 = new _dafny.Set();
  for (const _compr_44 of _dafny.IntegerRange(new BigNumber(349), new BigNumber(538))) {
    let _920_v127 = _compr_44;
    if (((new BigNumber(349)).isLessThanOrEqualTo(_920_v127)) && ((_920_v127).isLessThan(new BigNumber(538)))) {
      _coll44.add(_module.__default.safeDivisionInt(_920_v127, _918_p2));
    }
  }
  return _coll44;
}()).length)))).cardinality()), new BigNumber((_dafny.Seq.of(_917_v66)).length), _dafny.Seq.UnicodeFromString("qxnwwoae"), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length)), _918_p2);
            })(_841_v66, p2);
            let _nw165 = Array((new BigNumber(12)).toNumber());
            for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw165.length); _i0_30++) {
              _nw165[_i0_30] = _init30(new BigNumber(_i0_30));
            }
            _916_v128 = _nw165;
            let _921_v129;
            _921_v129 = _dafny.Map.Empty.slice().updateUnsafe(((_745_v0) ? (_915_v126) : (_915_v126)),_916_v128);
            _921_v129 = (_921_v129).update(_915_v126, _916_v128);
            let _922_v130;
            _922_v130 = _module.D11.create_DC30();
            let _923_v131;
            _923_v131 = _dafny.Seq.of(_module.D11.create_DC30(), _922_v130, _922_v130, _922_v130);
            let _924_v132;
            _924_v132 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
            let _925_v133;
            _925_v133 = _dafny.Seq.of(_dafny.Seq.of(p0, new BigNumber((_924_v132).cardinality())));
            let _926_v134;
            _926_v134 = _dafny.Seq.UnicodeFromString("sxyqh");
            let _927_v135;
            let _nw166 = Array((new BigNumber(19)).toNumber());
            _nw166[(_dafny.ZERO).toNumber()] = p0;
            _nw166[(_dafny.ONE).toNumber()] = p2;
            _nw166[(new BigNumber(2)).toNumber()] = ((_this).f12).multipliedBy(p3);
            _nw166[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(73)), function (_928_i17) {
              return _module.D11.create_DC30();
            }), _923_v131)).length));
            _nw166[(new BigNumber(4)).toNumber()] = p2;
            _nw166[(new BigNumber(5)).toNumber()] = p3;
            _nw166[(new BigNumber(6)).toNumber()] = p2;
            _nw166[(new BigNumber(7)).toNumber()] = p3;
            _nw166[(new BigNumber(8)).toNumber()] = ((_this).f12).plus(p3);
            _nw166[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_925_v133).length), p3);
            _nw166[(new BigNumber(10)).toNumber()] = (_this).f12;
            _nw166[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p3), new BigNumber((_dafny.Seq.of(_745_v0, !(_745_v0), !(_745_v0), _745_v0, _745_v0)).length));
            _nw166[(new BigNumber(12)).toNumber()] = new BigNumber(-586);
            _nw166[(new BigNumber(13)).toNumber()] = (_this).f12;
            _nw166[(new BigNumber(14)).toNumber()] = p2;
            _nw166[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt(p0, (_this).f12);
            _nw166[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_926_v134, _926_v134)).length);
            _nw166[(new BigNumber(17)).toNumber()] = p0;
            _nw166[(new BigNumber(18)).toNumber()] = p2;
            _927_v135 = _nw166;
            let _index114 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_927_v135).length));
            (_927_v135)[_index114] = (p2).plus(new BigNumber((_926_v134).length));
            let _index115 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_927_v135).length));
            (_927_v135)[_index115] = p2;
          }
        }
      }
      let _929_v136;
      let _nw167 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
      _929_v136 = _nw167;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_929_v136).length))) {
        let _930_i18 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_930_i18)) && ((_930_i18).isLessThan(new BigNumber((_929_v136).length))))) {
          (_929_v136)[(_930_i18)] = _dafny.Seq.Concat(_dafny.Seq.of(_745_v0, _745_v0), _dafny.Seq.Concat(_dafny.Seq.of(_745_v0, _745_v0, _745_v0), _dafny.Seq.of(false)));
        }
      }
      let _931_v137;
      _931_v137 = _dafny.Seq.UnicodeFromString("r");
      _931_v137 = _dafny.Seq.Concat(_931_v137, _931_v137);
      r0 = _745_v0;
      r1 = p3;
      let _932_v138;
      _932_v138 = _module.D11.create_DC31((_this).f12, (_dafny.ZERO).minus(p2));
      let _pat_let_tv26 = _745_v0;
      let _pat_let_tv27 = _745_v0;
      let _pat_let_tv28 = _745_v0;
      let _pat_let_tv29 = _745_v0;
      r2 = function (_source12) {
        if (_source12.is_DC30) {
          return _pat_let_tv26;
        } else if (_source12.is_DC31) {
          let _933___mcc_h3 = (_source12).cf44;
          let _934___mcc_h4 = (_source12).cf45;
          let _935_cf45 = _934___mcc_h4;
          let _936_cf44 = _933___mcc_h3;
          return _pat_let_tv27;
        } else if (_source12.is_DC29) {
          let _937___mcc_h5 = (_source12).cf43;
          let _938_cf43 = _937___mcc_h5;
          return _pat_let_tv28;
        } else {
          let _939___mcc_h6 = (_source12).cf46;
          let _940_cf46 = _939___mcc_h6;
          return _pat_let_tv29;
        }
      }(_932_v138);
      r3 = !(!(_745_v0) || (_745_v0));
      return [r0, r1, r2, r3];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      r0 = p3;
      (globalState).f2 = !(!(p1));
      let _941_v0;
      _941_v0 = _module.D0.create_DC1(p1);
      let _942_v1;
      _942_v1 = _module.D0.create_DC2(_941_v0);
      let _943_v2;
      _943_v2 = _module.D12.create_DC34(_942_v1, p1, true);
      let _944_i0;
      _944_i0 = _dafny.ZERO;
      L7: {
        while ((_943_v2).dtor_cf50) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_944_i0)) {
              break L7;
            }
            _944_i0 = (_944_i0).plus(_dafny.ONE);
            let _945_v3;
            let _nw168 = new _module.C2();
            _nw168.__ctor(p1);
            _945_v3 = _nw168;
            let _946_v4;
            _946_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,_945_v3);
            let _947_v7;
            _947_v7 = _module.D0.create_DC0(function () {
  let _coll45 = new _dafny.Map();
  for (const _compr_45 of (function () {
    let _coll46 = new _dafny.Map();
    for (const _compr_46 of _dafny.IntegerRange(new BigNumber(695), new BigNumber(246))) {
      let _948_v6 = _compr_46;
      if (((new BigNumber(695)).isLessThanOrEqualTo(_948_v6)) && ((_948_v6).isLessThan(new BigNumber(246)))) {
        _coll46.push([_module.__default.safeModuloInt(_948_v6, (_dafny.ZERO).minus(p3)),(_945_v3).f13]);
      }
    }
    return _coll46;
  }()).Keys.Elements) {
    let _949_v5 = _compr_45;
    if ((function () {
      let _coll47 = new _dafny.Map();
      for (const _compr_47 of _dafny.IntegerRange(new BigNumber(695), new BigNumber(246))) {
        let _948_v6 = _compr_47;
        if (((new BigNumber(695)).isLessThanOrEqualTo(_948_v6)) && ((_948_v6).isLessThan(new BigNumber(246)))) {
          _coll47.push([_module.__default.safeModuloInt(_948_v6, (_dafny.ZERO).minus(p3)),(_945_v3).f13]);
        }
      }
      return _coll47;
    }()).contains(_949_v5)) {
      _coll45.push([_module.__default.safeDivisionInt(_949_v5, p3),(_945_v3).f13]);
    }
  }
  return _coll45;
}());
            let _950_v8;
            _950_v8 = _dafny.MultiSet.fromElements((_945_v3).f13);
            let _951_v9;
            _951_v9 = _dafny.Set.fromElements(p1);
            let _952_v10;
            let _nw169 = Array((new BigNumber(17)).toNumber());
            _nw169[(_dafny.ZERO).toNumber()] = !(_946_v4).contains(p1);
            _nw169[(_dafny.ONE).toNumber()] = (_945_v3).f13;
            _nw169[(new BigNumber(2)).toNumber()] = true;
            _nw169[(new BigNumber(3)).toNumber()] = p1;
            _nw169[(new BigNumber(4)).toNumber()] = ((_945_v3).f13) && (p1);
            _nw169[(new BigNumber(5)).toNumber()] = p1;
            _nw169[(new BigNumber(6)).toNumber()] = ((_this).f12).isLessThanOrEqualTo(p3);
            _nw169[(new BigNumber(7)).toNumber()] = !(p1) || (!((_945_v3).f13));
            _nw169[(new BigNumber(8)).toNumber()] = (_945_v3).f13;
            _nw169[(new BigNumber(9)).toNumber()] = ((_945_v3).f13) === (p1);
            _nw169[(new BigNumber(10)).toNumber()] = _module.__default.fm1((_945_v3).f13, globalState);
            _nw169[(new BigNumber(11)).toNumber()] = (_945_v3).f13;
            _nw169[(new BigNumber(12)).toNumber()] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-224)), function (_953_i1) {
              return new _dafny.CodePoint('k'.codePointAt(0));
            }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), function (_954_i2) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            }));
            _nw169[(new BigNumber(13)).toNumber()] = (((_945_v3).f13) ? (p1) : ((_945_v3).f13));
            _nw169[(new BigNumber(14)).toNumber()] = (_this).fm4(_947_v7, _947_v7, _950_v8, globalState);
            _nw169[(new BigNumber(15)).toNumber()] = ((_this).f12).isEqualTo(new BigNumber((_951_v9).length));
            _nw169[(new BigNumber(16)).toNumber()] = (_945_v3).f13;
            _952_v10 = _nw169;
            let _955_v11;
            _955_v11 = _dafny.Seq.UnicodeFromString("akjkn");
            let _index116 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_952_v10).length));
            (_952_v10)[_index116] = (_945_v3).f13;
            let _956_v12;
            _956_v12 = _dafny.MultiSet.fromElements(_942_v1);
            let _index117 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_952_v10).length));
            let _rhs131 = _952_v10;
            let _rhs132 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), ((_957_p1, _958_v3) => function (_959_i3) {
              return (_module.D10.create_DC27(_957_p1, (_this).f12, new _dafny.CodePoint('n'.codePointAt(0)), _dafny.MultiSet.fromElements(false, (_958_v3).f13, false, _957_p1))).dtor_cf40;
            })(p1, _945_v3));
            let _rhs133 = (_956_v12).IsDisjointFrom((_956_v12).Intersect(_module.__default.fm31(globalState)));
            let _lhs98 = _952_v10;
            let _lhs99 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_952_v10).length));
            _952_v10 = _rhs131;
            _955_v11 = _rhs132;
            _lhs98[_lhs99] = _rhs133;
            (globalState).f5 = (_this).f12;
            let _960_v13;
            _960_v13 = new _dafny.CodePoint('s'.codePointAt(0));
            _960_v13 = new _dafny.CodePoint('f'.codePointAt(0));
            let _961_v14;
            _961_v14 = _dafny.Set.fromElements((_this).f12, (_this).f12);
            let _962_v15;
            _962_v15 = _module.D4.create_DC13(p3, _952_v10);
            let _963_v16;
            let _nw170 = new _module.C0();
            _nw170.__ctor((_961_v14).equals(_961_v14), (_962_v15).dtor_cf20);
            _963_v16 = _nw170;
          }
        }
      }
      let _964_v17;
      _964_v17 = _dafny.MultiSet.fromElements(p3, (_this).f12, p3);
      let _965_v18;
      _965_v18 = _dafny.Seq.of(p1, p1, p1, !(p1), _module.__default.fm1(!(p1), globalState));
      let _hi4 = (_this).f12;
      for (let _966_i4 = (new BigNumber((_964_v17).cardinality())).plus(new BigNumber((_965_v18).length)); _966_i4.isLessThan(_hi4); _966_i4 = _966_i4.plus(_dafny.ONE)) {
        _965_v18 = _965_v18;
        let _967_v19;
        let _nw171 = new _module.C2();
        _nw171.__ctor(p1);
        _967_v19 = _nw171;
        let _968_v20;
        let _nw172 = Array((new BigNumber(21)).toNumber()).fill(false);
        _968_v20 = _nw172;
        let _969_v21;
        let _nw173 = new _module.C0();
        _nw173.__ctor((_965_v18)[_module.__default.safeIndex(p3, new BigNumber((_965_v18).length))], _968_v20);
        _969_v21 = _nw173;
        let _970_v22;
        _970_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),(_dafny.ZERO).minus(p3));
        _970_v22 = (_970_v22).update((_967_v19).f13, (_this).f12);
      }
      let _hi5 = ((_this).f12).multipliedBy(p3);
      for (let _971_i5 = _module.__default.safeDivisionInt(new BigNumber((_964_v17).cardinality()), p3); _971_i5.isLessThan(_hi5); _971_i5 = _971_i5.plus(_dafny.ONE)) {
        (globalState).f5 = new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("wjl"))).length);
        let _972_v23;
        _972_v23 = _module.D7.create_DC18();
        let _973_v24;
        _973_v24 = _module.D7.create_DC19(_972_v23);
        let _974_v25;
        _974_v25 = _module.D7.create_DC19(_973_v24);
        let _975_v26;
        _975_v26 = _module.D7.create_DC19((_974_v25).dtor_cf24);
        let _976_v27;
        _976_v27 = _module.D7.create_DC19(_972_v23);
        let _977_v28;
        _977_v28 = _dafny.Map.Empty.slice().updateUnsafe(_974_v25,true);
        _977_v28 = (_977_v28).update(_976_v27, p1);
        let _978_v29;
        let _nw174 = Array((new BigNumber(19)).toNumber()).fill(false);
        _978_v29 = _nw174;
        let _979_v30;
        let _nw175 = new _module.C0();
        _nw175.__ctor(!(p1) || (p1), _978_v29);
        _979_v30 = _nw175;
        (globalState).f2 = !((new BigNumber((_965_v18).length)).isLessThanOrEqualTo(new BigNumber((p0).length)));
      }
      let _980_v31;
      let _nw176 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _980_v31 = _nw176;
      _980_v31 = _980_v31;
      let _981_v32;
      _981_v32 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
      let _982_v33;
      _982_v33 = _dafny.Map.Empty.slice().updateUnsafe(_981_v32,(_this).f12);
      r0 = (_module.D2.create_DC7((_this).f12, _982_v33, p3)).dtor_cf9;
      let _983_v34;
      let _nw177 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _983_v34 = _nw177;
      r1 = _983_v34;
      return [r0, r1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f9) {
      let _this = this;
      (_this)._f9 = f9;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber(((_dafny.MultiSet.fromElements(true)).Intersect(_dafny.MultiSet.fromElements(false, true))).cardinality())).isLessThan(_module.__default.safeModuloInt((_this).f9, (_this).f9));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((new BigNumber((_dafny.Seq.UnicodeFromString("urcwm")).length)).isEqualTo((_this).f9), _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("dwdsvbe"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-918)), function (_984_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })), false)).length));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      r1 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((p0).multipliedBy(p2)), p3);
      let _985_v0;
      let _nw178 = Array((new BigNumber(6)).toNumber()).fill(false);
      _985_v0 = _nw178;
      let _986_v1;
      let _nw179 = new _module.C0();
      _nw179.__ctor(true, _985_v0);
      _986_v1 = _nw179;
      let _987_v2;
      let _nw180 = new _module.C0();
      _nw180.__ctor(((_986_v1).f10) === ((_986_v1).f10), _985_v0);
      _987_v2 = _nw180;
      let _index118 = _module.__default.safeIndex(new BigNumber(204), new BigNumber(((_986_v1).f11).length));
      ((_986_v1).f11)[_index118] = (_986_v1).f10;
      let _988_v3;
      _988_v3 = new _dafny.CodePoint('m'.codePointAt(0));
      let _989_v4;
      let _init31 = ((_990_p0) => function (_991_i0) {
        return _module.__default.safeDivisionInt(_991_i0, _990_p0);
      })(p0);
      let _nw181 = Array((new BigNumber(8)).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw181.length); _i0_31++) {
        _nw181[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _989_v4 = _nw181;
      let _index119 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_989_v4).length));
      (_989_v4)[_index119] = p0;
      let _992_v5;
      _992_v5 = _dafny.MultiSet.fromElements((_986_v1).f10);
      let _index120 = _module.__default.safeIndex(new BigNumber(204), new BigNumber(((_986_v1).f11).length));
      let _index121 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_989_v4).length));
      let _rhs134 = !((_992_v5).Intersect(_dafny.MultiSet.fromElements((_986_v1).f10, (_986_v1).f10))).equals((_992_v5).Difference(_992_v5));
      let _rhs135 = _988_v3;
      let _rhs136 = _module.__default.safeDivisionInt(p2, new BigNumber((_992_v5).cardinality()));
      let _lhs100 = (_986_v1).f11;
      let _lhs101 = _module.__default.safeIndex(new BigNumber(204), new BigNumber(((_986_v1).f11).length));
      let _lhs102 = _989_v4;
      let _lhs103 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_989_v4).length));
      _lhs100[_lhs101] = _rhs134;
      _988_v3 = _rhs135;
      _lhs102[_lhs103] = _rhs136;
      let _993_v6;
      _993_v6 = _dafny.Seq.of((_987_v2).f10, (_986_v1).f10, (_987_v2).f10);
      let _994_v7;
      _994_v7 = _dafny.Map.Empty.slice().updateUnsafe((_993_v6)[_module.__default.safeIndex(p2, new BigNumber((_993_v6).length))],(_989_v4)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_989_v4).length))]);
      _994_v7 = (_994_v7).update(((_986_v1).f11)[_module.__default.safeIndex(new BigNumber(204), new BigNumber(((_986_v1).f11).length))], (_989_v4)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_989_v4).length))]);
      r2 = ((_986_v1).f11)[_module.__default.safeIndex(new BigNumber(204), new BigNumber(((_986_v1).f11).length))];
      r0 = (_986_v1).f10;
      r1 = p2;
      r2 = _module.__default.fm1((_987_v2).f10, globalState);
      let _995_v8;
      _995_v8 = _module.D0.create_DC1((_987_v2).f10);
      let _pat_let_tv30 = _987_v2;
      let _pat_let_tv31 = _986_v1;
      let _pat_let_tv32 = _986_v1;
      r3 = function (_source13) {
        if (_source13.is_DC1) {
          let _996___mcc_h0 = (_source13).cf1;
          let _997_cf1 = _996___mcc_h0;
          return (_dafny.Set.fromElements(false)).IsDisjointFrom(_dafny.Set.fromElements(true));
        } else if (_source13.is_DC0) {
          let _998___mcc_h1 = (_source13).cf0;
          let _999_cf0 = _998___mcc_h1;
          return !(false) || ((_pat_let_tv30).f10);
        } else {
          let _1000___mcc_h2 = (_source13).cf2;
          let _1001_cf2 = _1000___mcc_h2;
          return ((_pat_let_tv31).f10) && ((_pat_let_tv32).f10);
        }
      }(_995_v8);
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1002_v0;
      _1002_v0 = _dafny.Seq.UnicodeFromString("vg");
      let _1003_v1;
      _1003_v1 = false;
      let _1004_v2;
      _1004_v2 = _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1002_v0).length),_1003_v1));
      let _pat_let_tv33 = _1003_v1;
      let _pat_let_tv34 = _1003_v1;
      let _pat_let_tv35 = _1003_v1;
      let _source14 = function (_source15) {
        if (_source15.is_DC1) {
          let _1005___mcc_h3 = (_source15).cf1;
          let _1006_cf1 = _1005___mcc_h3;
          if (true) {
            return _module.D0.create_DC1(true);
          } else {
            return _module.D0.create_DC1(_pat_let_tv33);
          }
        } else if (_source15.is_DC0) {
          let _1007___mcc_h4 = (_source15).cf0;
          let _1008_cf0 = _1007___mcc_h4;
          return _module.D0.create_DC1(_pat_let_tv34);
        } else {
          let _1009___mcc_h5 = (_source15).cf2;
          let _1010_cf2 = _1009___mcc_h5;
          return _module.D0.create_DC1(!(_pat_let_tv35));
        }
      }(_1004_v2);
      if (_source14.is_DC1) {
        let _1011___mcc_h0 = (_source14).cf1;
        let _1012_cf1 = _1011___mcc_h0;
        r1 = p0;
        (globalState).f5 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        let _1013_v3;
        _1013_v3 = _dafny.Seq.of(false);
        let _1014_v4;
        _1014_v4 = _dafny.Seq.of((p1).isLessThan(new BigNumber((_1013_v3).length)));
        let _rhs137 = _1012_cf1;
        let _rhs138 = _1014_v4;
        let _rhs139 = new BigNumber((_1014_v4).length);
        let _lhs104 = globalState;
        let _lhs105 = globalState;
        _lhs104.f0 = _rhs137;
        _1013_v3 = _rhs138;
        _lhs105.f5 = _rhs139;
        let _1015_v5;
        _1015_v5 = new _dafny.CodePoint('c'.codePointAt(0));
        _1015_v5 = _1015_v5;
      } else if (_source14.is_DC0) {
        let _1016___mcc_h1 = (_source14).cf0;
        let _1017_cf0 = _1016___mcc_h1;
        let _1018_v6;
        _1018_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1003_v1,p1);
        let _1019_v7;
        let _nw182 = Array((new BigNumber(26)).toNumber());
        _nw182[(_dafny.ZERO).toNumber()] = true;
        _nw182[(_dafny.ONE).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(2)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(3)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(4)).toNumber()] = !(_1003_v1);
        _nw182[(new BigNumber(5)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(6)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(7)).toNumber()] = !(_1003_v1);
        _nw182[(new BigNumber(8)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(9)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(10)).toNumber()] = false;
        _nw182[(new BigNumber(11)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(12)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(13)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(14)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(15)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(16)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(17)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(18)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(19)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(20)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(21)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(22)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(23)).toNumber()] = !(_1003_v1);
        _nw182[(new BigNumber(24)).toNumber()] = _1003_v1;
        _nw182[(new BigNumber(25)).toNumber()] = _1003_v1;
        _1019_v7 = _nw182;
        let _1020_v8;
        let _nw183 = new _module.C0();
        _nw183.__ctor(_1003_v1, _1019_v7);
        _1020_v8 = _nw183;
        let _1021_v9;
        _1021_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1018_v6,_1020_v8);
        let _1022_v10;
        _1022_v10 = _module.D0.create_DC0(_1017_cf0);
        let _1023_v11;
        _1023_v11 = _module.D0.create_DC2(_1022_v10);
        let _rhs140 = ((_1021_v9).Merge(_1021_v9)).update((_1018_v6).Merge(_1018_v6), ((_1003_v1) ? (_1020_v8) : (_1020_v8)));
        let _rhs141 = _1023_v11;
        _1021_v9 = _rhs140;
        _1023_v11 = _rhs141;
        if (_1003_v1) {
          _1017_cf0 = (_1017_cf0).update(p1, _1003_v1);
          let _1024_v12;
          _1024_v12 = _dafny.Set.fromElements(false);
          let _1025_v13;
          _1025_v13 = _dafny.Seq.of(_1024_v12);
          let _1026_v14;
          _1026_v14 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1027_v15;
          _1027_v15 = _dafny.MultiSet.fromElements((_module.D2.create_DC5(_1020_v8)).dtor_cf5, _1020_v8, _1020_v8);
          (globalState).f5 = _module.__default.fm0((_this).f9, new BigNumber((_1002_v0).length), (p1).plus((_this).fm5(_1025_v13, _1026_v14, _dafny.Seq.of(new BigNumber((_1027_v15).cardinality())), _dafny.Seq.of(new BigNumber(835)), globalState)), _1003_v1, globalState);
          let _1028_v16;
          _1028_v16 = _dafny.MultiSet.fromElements(_1003_v1);
          let _1029_v17;
          _1029_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1028_v16,_1003_v1);
          _1026_v14 = _module.__default.fm9((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((p2).length), new BigNumber((_1029_v17).length))), globalState);
          (globalState).f5 = p1;
          _1017_cf0 = (_1017_cf0).update((_dafny.ZERO).minus(p1), true);
        } else {
          let _1030_v18;
          let _init32 = function (_1031_i0) {
            return _module.__default.safeDivisionInt(_1031_i0, (_this).f9);
          };
          let _nw184 = Array((new BigNumber(21)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw184.length); _i0_32++) {
            _nw184[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1030_v18 = _nw184;
          let _1032_v19;
          _1032_v19 = _dafny.Set.fromElements(new BigNumber((_1002_v0).length));
          let _rhs142 = _1030_v18;
          let _rhs143 = false;
          let _rhs144 = _1032_v19;
          let _lhs106 = globalState;
          _1030_v18 = _rhs142;
          _lhs106.f0 = _rhs143;
          _1032_v19 = _rhs144;
          let _1033_v20;
          _1033_v20 = _dafny.MultiSet.fromElements(_1003_v1);
          (globalState).f5 = _module.__default.safeDivisionInt(p1, (((_1033_v20).contains((_1020_v8).fm4(_1004_v2, _1004_v2, _1033_v20, globalState))) ? ((_1033_v20).get((_1020_v8).fm4(_1004_v2, _1004_v2, _1033_v20, globalState))) : (new BigNumber(-158))));
          let _rhs145 = (p1).plus((_this).f9);
          let _rhs146 = (_this).f9;
          let _rhs147 = p1;
          let _lhs107 = globalState;
          r1 = _rhs145;
          _lhs107.f5 = _rhs146;
          r1 = _rhs147;
          let _1034_v21;
          _1034_v21 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1035_v22;
          _1035_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1003_v1,_1034_v21);
          let _1036_v23;
          _1036_v23 = _module.D3.create_DC8(_1002_v0);
          let _1037_v24;
          let _nw185 = Array((new BigNumber(21)).toNumber());
          _nw185[(_dafny.ZERO).toNumber()] = _1002_v0;
          _nw185[(_dafny.ONE).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(2)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1002_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(667)), function (_1038_i1) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }));
          _nw185[(new BigNumber(4)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1002_v0, _1002_v0);
          _nw185[(new BigNumber(6)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1002_v0, _1002_v0);
          _nw185[(new BigNumber(8)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1002_v0, _1002_v0);
          _nw185[(new BigNumber(10)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(11)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(12)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1002_v0, _1002_v0);
          _nw185[(new BigNumber(14)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-366)), function (_1039_i2) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          });
          _nw185[(new BigNumber(16)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm8(_1003_v1, _dafny.Map.Empty.slice().updateUnsafe(true,_1034_v21), _1003_v1, globalState), _1002_v0);
          _nw185[(new BigNumber(18)).toNumber()] = _module.__default.fm8(_1003_v1, _1035_v22, _1003_v1, globalState);
          _nw185[(new BigNumber(19)).toNumber()] = _1002_v0;
          _nw185[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(273)), ((_1040_v21) => function (_1041_i3) {
            return _1040_v21;
          })(_1034_v21)), (_1036_v23).dtor_cf10);
          _1037_v24 = _nw185;
          let _index122 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_1037_v24).length));
          (_1037_v24)[_index122] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ggh"), _1002_v0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ggh"), _1002_v0)).length)), _1034_v21), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ggh"), _1002_v0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ggh"), _1002_v0)).length)), _1034_v21)).length)), _1034_v21);
          let _index123 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_1037_v24).length));
          (_1037_v24)[_index123] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(827)), ((_1042_v21) => function (_1043_i4) {
            return _1042_v21;
          })(_1034_v21)), _1002_v0);
          let _1044_v25;
          let _nw186 = Array((new BigNumber(28)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw186[(_dafny.ONE).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(2)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
          _nw186[(new BigNumber(4)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(5)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _nw186[(new BigNumber(7)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(8)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(9)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(10)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(11)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(12)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(13)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
          _nw186[(new BigNumber(15)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(16)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(17)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(18)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(19)).toNumber()] = ((_1037_v24)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_1037_v24).length))])[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1033_v20).cardinality())), new BigNumber(((_1037_v24)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_1037_v24).length))]).length))];
          _nw186[(new BigNumber(20)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(21)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(22)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(23)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(24)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(25)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(26)).toNumber()] = _1034_v21;
          _nw186[(new BigNumber(27)).toNumber()] = _1034_v21;
          _1044_v25 = _nw186;
          let _index124 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1044_v25).length));
          (_1044_v25)[_index124] = _1034_v21;
          let _index125 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1044_v25).length));
          (_1044_v25)[_index125] = _1034_v21;
          let _index126 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1044_v25).length));
          let _index127 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1044_v25).length));
          let _rhs148 = _1030_v18;
          let _rhs149 = (p0).minus(p0);
          let _rhs150 = _1034_v21;
          let _rhs151 = _1034_v21;
          let _lhs108 = globalState;
          let _lhs109 = _1044_v25;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1044_v25).length));
          let _lhs111 = _1044_v25;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1044_v25).length));
          _1030_v18 = _rhs148;
          _lhs108.f5 = _rhs149;
          _lhs109[_lhs110] = _rhs150;
          _lhs111[_lhs112] = _rhs151;
        }
        let _1045_v26;
        let _nw187 = Array((new BigNumber(23)).toNumber()).fill([]);
        _1045_v26 = _nw187;
        let _1046_v27;
        let _init33 = ((_1047_cf0, _1048_v0) => function (_1049_i5) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_module.D3.create_DC10(_module.D3.create_DC8(_dafny.Seq.UnicodeFromString("ewwnswuai"))), _module.D3.create_DC10(_module.D3.create_DC9(new BigNumber((_1047_cf0).length)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(724)), ((_1050_v0) => function (_1051_i6) {
            return _module.D3.create_DC10((_module.D3.create_DC10(_module.D3.create_DC8(_1050_v0))).dtor_cf12);
          })(_1048_v0)));
        })(_1017_cf0, _1002_v0);
        let _nw188 = Array((new BigNumber(6)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw188.length); _i0_33++) {
          _nw188[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1046_v27 = _nw188;
        let _1052_v28;
        _1052_v28 = _module.D3.create_DC8(_1002_v0);
        let _1053_v29;
        _1053_v29 = _dafny.Seq.of(_module.D3.create_DC10(_1052_v28));
        let _index128 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1046_v27).length));
        (_1046_v27)[_index128] = _dafny.Seq.Concat(_1053_v29, _1053_v29);
        let _1054_v30;
        _1054_v30 = _dafny.Seq.of((true) || (_1003_v1));
        let _index129 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1046_v27).length));
        let _rhs152 = (_1054_v30)[_module.__default.safeIndex(p1, new BigNumber((_1054_v30).length))];
        let _rhs153 = _1045_v26;
        let _rhs154 = _1053_v29;
        let _rhs155 = (_this).f9;
        let _rhs156 = ((_this).f9).isLessThan(new BigNumber((_dafny.Seq.Concat(_1054_v30, _dafny.Seq.of(_1003_v1))).length));
        let _lhs113 = globalState;
        let _lhs114 = _1046_v27;
        let _lhs115 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1046_v27).length));
        let _lhs116 = globalState;
        _lhs113.f2 = _rhs152;
        _1045_v26 = _rhs153;
        _lhs114[_lhs115] = _rhs154;
        _lhs116.f5 = _rhs155;
        r2 = _rhs156;
        (globalState).f5 = p0;
      } else {
        let _1055___mcc_h2 = (_source14).cf2;
        let _1056_cf2 = _1055___mcc_h2;
        let _rhs157 = (_module.__default.safeDivisionInt(p0, (_this).f9)).minus((_this).f9);
        let _rhs158 = _1003_v1;
        let _rhs159 = (_this).f9;
        let _lhs117 = globalState;
        let _lhs118 = globalState;
        _lhs117.f5 = _rhs157;
        _1003_v1 = _rhs158;
        _lhs118.f5 = _rhs159;
        let _1057_v31;
        let _nw189 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1057_v31 = _nw189;
        let _index130 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_1057_v31).length));
        (_1057_v31)[_index130] = new BigNumber(903);
        let _index131 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_1057_v31).length));
        (_1057_v31)[_index131] = _module.__default.safeDivisionInt(p1, (p0).plus(new BigNumber(717)));
        let _1058_v32;
        _1058_v32 = _dafny.Seq.of(_1057_v31);
        let _1059_v33;
        let _nw190 = Array((new BigNumber(2)).toNumber());
        _nw190[(_dafny.ZERO).toNumber()] = (_1058_v32)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("sbtwqs")).length), new BigNumber((_1058_v32).length))];
        _nw190[(_dafny.ONE).toNumber()] = _1057_v31;
        _1059_v33 = _nw190;
        let _1060_v34;
        let _nw191 = Array((new BigNumber(21)).toNumber()).fill([]);
        _1060_v34 = _nw191;
        let _1061_v35;
        let _nw192 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Set.Empty);
        _1061_v35 = _nw192;
        let _index132 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1060_v34).length));
        (_1060_v34)[_index132] = _1061_v35;
        let _index133 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1060_v34).length));
        let _rhs160 = _1059_v33;
        let _rhs161 = _1061_v35;
        let _rhs162 = p1;
        let _rhs163 = _1002_v0;
        let _rhs164 = _1003_v1;
        let _lhs119 = _1060_v34;
        let _lhs120 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1060_v34).length));
        let _lhs121 = globalState;
        let _lhs122 = globalState;
        _1059_v33 = _rhs160;
        _lhs119[_lhs120] = _rhs161;
        _lhs121.f5 = _rhs162;
        _1002_v0 = _rhs163;
        _lhs122.f2 = _rhs164;
        let _1062_v36;
        _1062_v36 = _module.D3.create_DC9(p0);
        let _pat_let_tv36 = _1062_v36;
        let _1063_v37;
        _1063_v37 = _dafny.MultiSet.fromElements(_module.D3.create_DC10(_1062_v36), function (_pat_let16_0) {
          return function (_1064_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_1065_dt__update_hcf12_h0) {
                return _module.D3.create_DC10(_1065_dt__update_hcf12_h0);
              }(_pat_let17_0);
            }(_pat_let_tv36);
          }(_pat_let16_0);
        }(_module.D3.create_DC10(_1062_v36)));
        let _1066_v38;
        _1066_v38 = _module.D3.create_DC10(_1062_v36);
        let _1067_v39;
        _1067_v39 = _module.D0.create_DC1(_1003_v1);
        let _1068_v40;
        let _nw193 = Array((new BigNumber(15)).toNumber());
        _nw193[(_dafny.ZERO).toNumber()] = _1003_v1;
        _nw193[(_dafny.ONE).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(2)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(3)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(4)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(5)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(6)).toNumber()] = true;
        _nw193[(new BigNumber(7)).toNumber()] = true;
        _nw193[(new BigNumber(8)).toNumber()] = true;
        _nw193[(new BigNumber(9)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(10)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(11)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(12)).toNumber()] = _1003_v1;
        _nw193[(new BigNumber(13)).toNumber()] = true;
        _nw193[(new BigNumber(14)).toNumber()] = (_1067_v39).dtor_cf1;
        _1068_v40 = _nw193;
        let _pat_let_tv37 = _1062_v36;
        let _1069_v41;
        let _nw194 = new _module.C0();
        _nw194.__ctor(!(((_1063_v37).update(function (_pat_let18_0) {
          return function (_1070_dt__update__tmp_h1) {
            return function (_pat_let19_0) {
              return function (_1071_dt__update_hcf12_h1) {
                return _module.D3.create_DC10(_1071_dt__update_hcf12_h1);
              }(_pat_let19_0);
            }(_pat_let_tv37);
          }(_pat_let18_0);
        }(_1066_v38), _module.__default.abs((_1057_v31)[_module.__default.safeIndex(new BigNumber(555), new BigNumber((_1057_v31).length))]))).IsSubsetOf(_1063_v37)), _1068_v40);
        _1069_v41 = _nw194;
        _1069_v41 = _1069_v41;
      }
      if (false) {
        let _1072_v42;
        _1072_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1003_v1);
        let _1073_v43;
        _1073_v43 = _dafny.Seq.of(_1072_v42);
        let _1074_v44;
        _1074_v44 = _dafny.MultiSet.fromElements(_1003_v1);
        let _1075_v45;
        _1075_v45 = _dafny.Set.fromElements(_1003_v1, _1003_v1);
        let _1076_v46;
        _1076_v46 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1077_v47;
        _1077_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1003_v1,_1003_v1);
        let _1078_v48;
        _1078_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1077_v47,new BigNumber((_1075_v45).length));
        let _1079_v49;
        _1079_v49 = _module.D2.create_DC7(p0, _1078_v48, new BigNumber((_dafny.Seq.update(_1002_v0, _module.__default.safeIndex(p1, new BigNumber((_1002_v0).length)), _1076_v46)).length));
        let _1080_v50;
        _1080_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1074_v44,_module.__default.fm10(_1075_v45, _1076_v46, _1079_v49, p0, globalState));
        let _rhs165 = _1003_v1;
        let _rhs166 = (_1080_v50).contains(_1074_v44);
        let _rhs167 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_1072_v42).update((_this).f9, _1003_v1), _dafny.Map.Empty.slice().updateUnsafe(p1,_1003_v1)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-788)), ((_1081_v42) => function (_1082_i7) {
          return _1081_v42;
        })(_1072_v42))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), ((_1083_v42) => function (_1084_i8) {
          return _1083_v42;
        })(_1072_v42)));
        _1003_v1 = _rhs165;
        _1003_v1 = _rhs166;
        _1073_v43 = _rhs167;
        (globalState).f0 = _1003_v1;
        let _1085_v51;
        _1085_v51 = _module.D2.create_DC6(_1003_v1);
        let _source16 = _1085_v51;
        if (_source16.is_DC6) {
          let _1086___mcc_h6 = (_source16).cf6;
          let _1087_cf6 = _1086___mcc_h6;
          let _1088_v52;
          let _init34 = ((_1089_cf6) => function (_1090_i9) {
            return _1089_cf6;
          })(_1087_cf6);
          let _nw195 = Array((new BigNumber(15)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw195.length); _i0_34++) {
            _nw195[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1088_v52 = _nw195;
          let _1091_v53;
          let _nw196 = new _module.C0();
          _nw196.__ctor(_1087_cf6, _1088_v52);
          _1091_v53 = _nw196;
          let _1092_v54;
          _1092_v54 = _dafny.Seq.of(_1091_v53, _1091_v53);
          let _1093_v55;
          let _nw197 = Array((new BigNumber(15)).toNumber());
          _nw197[(_dafny.ZERO).toNumber()] = _1087_cf6;
          _nw197[(_dafny.ONE).toNumber()] = _1003_v1;
          _nw197[(new BigNumber(2)).toNumber()] = !(true) || (_1003_v1);
          _nw197[(new BigNumber(3)).toNumber()] = _1087_cf6;
          _nw197[(new BigNumber(4)).toNumber()] = !_dafny.areEqual(_1092_v54, _dafny.Seq.of(_1091_v53));
          _nw197[(new BigNumber(5)).toNumber()] = _1003_v1;
          _nw197[(new BigNumber(6)).toNumber()] = _1087_cf6;
          _nw197[(new BigNumber(7)).toNumber()] = (_1091_v53).f10;
          _nw197[(new BigNumber(8)).toNumber()] = !(_1003_v1) || (_1087_cf6);
          _nw197[(new BigNumber(9)).toNumber()] = !_dafny.areEqual(_1002_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(130)), ((_1094_v46) => function (_1095_i10) {
            return _1094_v46;
          })(_1076_v46)));
          _nw197[(new BigNumber(10)).toNumber()] = _1087_cf6;
          _nw197[(new BigNumber(11)).toNumber()] = _1003_v1;
          _nw197[(new BigNumber(12)).toNumber()] = _1003_v1;
          _nw197[(new BigNumber(13)).toNumber()] = _1087_cf6;
          _nw197[(new BigNumber(14)).toNumber()] = _1003_v1;
          _1093_v55 = _nw197;
          let _index134 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_1093_v55).length));
          (_1093_v55)[_index134] = !(_1087_cf6);
          let _index135 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_1093_v55).length));
          (_1093_v55)[_index135] = (p0).isLessThan((_this).f9);
          (globalState).f2 = (_1091_v53).f10;
          let _1096_v56;
          let _nw198 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _1096_v56 = _nw198;
          _1096_v56 = ((_1087_cf6) ? (_1096_v56) : (_1096_v56));
          let _1097_v57;
          _1097_v57 = _dafny.Seq.of((_1093_v55)[_module.__default.safeIndex(new BigNumber(638), new BigNumber((_1093_v55).length))], (_1091_v53).f10);
          let _1098_v58;
          _1098_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1076_v46,(_1091_v53).f11);
          let _1099_v59;
          _1099_v59 = _module.D4.create_DC13(new BigNumber((_1097_v57).length), (((_1098_v58).contains(_1076_v46)) ? ((_1098_v58).get(_1076_v46)) : (_1093_v55)));
          let _1100_v60;
          _1100_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1099_v59).dtor_cf20,_1088_v52);
          let _1101_v61;
          let _nw199 = new _module.C0();
          _nw199.__ctor((_1091_v53).f10, (((_1100_v60).contains(_1088_v52)) ? ((_1100_v60).get(_1088_v52)) : (_1088_v52)));
          _1101_v61 = _nw199;
          let _1102_v62;
          _1102_v62 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1101_v61);
          let _1103_v63;
          _1103_v63 = _dafny.Seq.of(_1075_v45);
          let _1104_v64;
          _1104_v64 = _dafny.Seq.of(new BigNumber((_1097_v57).length), p1);
          let _1105_v65;
          _1105_v65 = _dafny.Set.fromElements(new BigNumber((_1097_v57).length));
          let _1106_v66;
          let _nw200 = Array((new BigNumber(16)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = _1102_v62;
          _nw200[(_dafny.ONE).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(2)).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(3)).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(4)).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(5)).toNumber()] = (_1102_v62).update((_this).f9, _1101_v61);
          _nw200[(new BigNumber(6)).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(7)).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_1101_v61);
          _nw200[(new BigNumber(9)).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false))).cardinality()),_1101_v61);
          _nw200[(new BigNumber(11)).toNumber()] = (_1102_v62).Merge(_1102_v62);
          _nw200[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1079_v49).dtor_cf7,_1101_v61);
          _nw200[(new BigNumber(13)).toNumber()] = (_1102_v62).update((_this).fm5(_1103_v63, _1076_v46, _1104_v64, _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1105_v65).length)), p0), globalState), _1101_v61);
          _nw200[(new BigNumber(14)).toNumber()] = _1102_v62;
          _nw200[(new BigNumber(15)).toNumber()] = (_1102_v62).Merge(_1102_v62);
          _1106_v66 = _nw200;
          let _index136 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_1106_v66).length));
          (_1106_v66)[_index136] = _1102_v62;
          let _index137 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_1106_v66).length));
          (_1106_v66)[_index137] = _1102_v62;
        } else if (_source16.is_DC7) {
          let _1107___mcc_h7 = (_source16).cf7;
          let _1108___mcc_h8 = (_source16).cf8;
          let _1109___mcc_h9 = (_source16).cf9;
          let _1110_cf9 = _1109___mcc_h9;
          let _1111_cf8 = _1108___mcc_h8;
          let _1112_cf7 = _1107___mcc_h7;
          (globalState).f0 = (((_1003_v1) ? (_1003_v1) : (_1003_v1))) && (_1003_v1);
          let _1113_v67;
          _1113_v67 = _dafny.Seq.of(_1112_cf7, new BigNumber((_1077_v47).length));
          _1113_v67 = _1113_v67;
          _1079_v49 = _1079_v49;
          let _1114_v68;
          _1114_v68 = _module.D0.create_DC1(false);
          let _1115_v69;
          _1115_v69 = _dafny.Seq.of(_1003_v1, _1003_v1, _1003_v1, _1003_v1, (_1114_v68).dtor_cf1);
          let _1116_v70;
          let _nw201 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1116_v70 = _nw201;
          let _1117_v71;
          _1117_v71 = _module.D4.create_DC13(_1110_cf9, _1116_v70);
          (globalState).f0 = (_1115_v69)[_module.__default.safeIndex((_1117_v71).dtor_cf19, new BigNumber((_1115_v69).length))];
        } else {
          let _1118___mcc_h10 = (_source16).cf5;
          let _1119_cf5 = _1118___mcc_h10;
          r2 = true;
          let _1120_v72;
          _1120_v72 = _dafny.Seq.of(_1003_v1);
          let _1121_v73;
          _1121_v73 = _module.D0.create_DC1((_1120_v72)[_module.__default.safeIndex(p0, new BigNumber((_1120_v72).length))]);
          let _1122_v74;
          _1122_v74 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC9(new BigNumber(610)),_1121_v73);
          _1122_v74 = _1122_v74;
          let _1123_v75;
          let _init35 = ((_1124_v1) => function (_1125_i11) {
            return _1124_v1;
          })(_1003_v1);
          let _nw202 = Array((new BigNumber(10)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw202.length); _i0_35++) {
            _nw202[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1123_v75 = _nw202;
          let _1126_v76;
          let _nw203 = new _module.C0();
          _nw203.__ctor(true, _1123_v75);
          _1126_v76 = _nw203;
          let _1127_v77;
          let _init36 = ((_1128_p1) => function (_1129_i12) {
            return (_1129_i12).plus(_1128_p1);
          })(p1);
          let _nw204 = Array((new BigNumber(10)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw204.length); _i0_36++) {
            _nw204[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1127_v77 = _nw204;
          let _1130_v78;
          _1130_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1126_v76,_1127_v77);
          let _1131_v79;
          _1131_v79 = _dafny.Map.Empty.slice().updateUnsafe((_1126_v76).f10,_1077_v47);
          let _1132_v80;
          _1132_v80 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_1131_v79).length), p0, p1, (_1126_v76).f10, globalState),p1);
          let _rhs168 = _1003_v1;
          let _rhs169 = ((_1003_v1) ? (new BigNumber(((_1130_v78).Merge(_1130_v78)).length)) : ((((_1126_v76).f10) ? ((_this).f9) : (p0))));
          let _rhs170 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,p0)).Merge((_1132_v80).Merge(_1132_v80))).length);
          let _lhs123 = globalState;
          let _lhs124 = globalState;
          _lhs123.f2 = _rhs168;
          r1 = _rhs169;
          _lhs124.f5 = _rhs170;
          let _1133_v81;
          _1133_v81 = _dafny.MultiSet.fromElements(_1072_v42, (_1072_v42).Merge(_1072_v42), _1072_v42, _1072_v42);
          let _index138 = _module.__default.safeIndex(new BigNumber(526), new BigNumber(((_1126_v76).f11).length));
          ((_1126_v76).f11)[_index138] = (_1126_v76).f10;
          let _1134_v82;
          _1134_v82 = _dafny.MultiSet.fromElements(_1002_v0);
          let _1135_v83;
          _1135_v83 = _dafny.Set.fromElements(p1);
          let _index139 = _module.__default.safeIndex(new BigNumber(526), new BigNumber(((_1126_v76).f11).length));
          let _rhs171 = p0;
          let _rhs172 = (new BigNumber((_1134_v82).cardinality())).multipliedBy((((_1126_v76).f10) ? ((_this).f9) : (_module.__default.fm0(new BigNumber((_1135_v83).length), p0, p1, (_1085_v51).dtor_cf6, globalState))));
          let _rhs173 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1120_v72, _dafny.Seq.of((_1126_v76).f10)), _1120_v72)).length);
          let _rhs174 = _1133_v81;
          let _rhs175 = _1003_v1;
          let _lhs125 = globalState;
          let _lhs126 = globalState;
          let _lhs127 = globalState;
          let _lhs128 = (_1126_v76).f11;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(526), new BigNumber(((_1126_v76).f11).length));
          _lhs125.f5 = _rhs171;
          _lhs126.f5 = _rhs172;
          _lhs127.f5 = _rhs173;
          _1133_v81 = _rhs174;
          _lhs128[_lhs129] = _rhs175;
        }
        if ((p1).isLessThanOrEqualTo(new BigNumber((_1075_v45).length))) {
          _1003_v1 = _1003_v1;
          let _1136_v84;
          _1136_v84 = _module.D0.create_DC1(_1003_v1);
          let _1137_v85;
          let _nw205 = Array((new BigNumber(3)).toNumber());
          _nw205[(_dafny.ZERO).toNumber()] = false;
          _nw205[(_dafny.ONE).toNumber()] = _1003_v1;
          _nw205[(new BigNumber(2)).toNumber()] = (_1136_v84).dtor_cf1;
          _1137_v85 = _nw205;
          let _1138_v86;
          let _nw206 = new _module.C0();
          _nw206.__ctor(!(!(((_this).f9).isLessThanOrEqualTo((_this).f9))), _1137_v85);
          _1138_v86 = _nw206;
          let _index140 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1137_v85).length));
          (_1137_v85)[_index140] = !(_1003_v1) || ((_1138_v86).f10);
          let _index141 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1137_v85).length));
          (_1137_v85)[_index141] = !(_1003_v1);
          let _1139_v87;
          _1139_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1074_v44,(_1137_v85)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_1137_v85).length))]);
          r1 = new BigNumber((((_1139_v87).Merge(_1139_v87)).Merge(_1139_v87)).length);
          let _1140_v88;
          let _nw207 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _1140_v88 = _nw207;
          let _1141_v89;
          let _nw208 = new _module.C2();
          _nw208.__ctor(!((_1138_v86).f10));
          _1141_v89 = _nw208;
          let _1142_v90;
          _1142_v90 = _dafny.Seq.of((_1137_v85)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_1137_v85).length))], (_1138_v86).f10, (_1137_v85)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_1137_v85).length))]);
          let _1143_v91;
          _1143_v91 = _module.D7.create_DC17(_1142_v90);
          let _1144_v92;
          _1144_v92 = _dafny.MultiSet.fromElements(_module.D7.create_DC17(_dafny.Seq.of((_1138_v86).f10)), _1143_v91);
          let _1145_v93;
          _1145_v93 = _dafny.MultiSet.fromElements(_1141_v89);
          let _rhs176 = _1140_v88;
          let _rhs177 = (_1137_v85)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_1137_v85).length))];
          let _rhs178 = (_1138_v86).f10;
          let _rhs179 = _1141_v89;
          let _rhs180 = _dafny.Set.fromElements(_1003_v1, (_1144_v92).IsDisjointFrom(_1144_v92), (_1145_v93).IsDisjointFrom(_1145_v93));
          let _lhs130 = globalState;
          let _lhs131 = globalState;
          _1140_v88 = _rhs176;
          _lhs130.f2 = _rhs177;
          r2 = _rhs178;
          _1141_v89 = _rhs179;
          _lhs131.f8 = _rhs180;
        } else {
          let _rhs181 = _1003_v1;
          let _rhs182 = p1;
          let _rhs183 = !(!(_1003_v1));
          let _rhs184 = (_this).f9;
          let _lhs132 = globalState;
          let _lhs133 = globalState;
          let _lhs134 = globalState;
          let _lhs135 = globalState;
          _lhs132.f2 = _rhs181;
          _lhs133.f5 = _rhs182;
          _lhs134.f2 = _rhs183;
          _lhs135.f5 = _rhs184;
          let _1146_v94;
          _1146_v94 = _dafny.Map.Empty.slice().updateUnsafe((p0).minus(new BigNumber(755)),(_dafny.ZERO).minus(p0));
          let _1147_v96;
          _1147_v96 = _dafny.Seq.of(_1003_v1, !(_1003_v1));
          _1146_v94 = (_1146_v94).update(_module.__default.fm0(p0, p1, new BigNumber((function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of (_dafny.Map.Empty.slice().updateUnsafe(p0,true)).Keys.Elements) {
              let _1148_v95 = _compr_48;
              if ((_dafny.Map.Empty.slice().updateUnsafe(p0,true)).contains(_1148_v95)) {
                _coll48.push([(_1148_v95).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p0, p0)).length))),new BigNumber((_1002_v0).length)]);
              }
            }
            return _coll48;
          }()).length), true, globalState), new BigNumber((_1147_v96).length));
          (globalState).f0 = true;
          let _1149_v97;
          let _nw209 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _1149_v97 = _nw209;
          let _index142 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1149_v97).length));
          (_1149_v97)[_index142] = p0;
          let _1150_v98;
          let _nw210 = Array((new BigNumber(14)).toNumber()).fill(false);
          _1150_v98 = _nw210;
          let _1151_v99;
          let _nw211 = new _module.C0();
          _nw211.__ctor(_1003_v1, _1150_v98);
          _1151_v99 = _nw211;
          let _1152_v100;
          _1152_v100 = _dafny.MultiSet.fromElements(_1151_v99, _1151_v99, _1151_v99, _1151_v99, _1151_v99);
          let _1153_v101;
          _1153_v101 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality()),(_this).f9)).length));
          let _1154_v102;
          _1154_v102 = _dafny.Seq.of(_1002_v0, _1002_v0);
          let _index143 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1149_v97).length));
          let _rhs185 = (((_1152_v100).contains(_1151_v99)) ? ((_1152_v100).get(_1151_v99)) : ((_dafny.ZERO).minus((((_1151_v99).f10) ? ((_1153_v101)[_module.__default.safeIndex(new BigNumber((_1147_v96).length), new BigNumber((_1153_v101).length))]) : ((_this).f9)))));
          let _rhs186 = (_dafny.ZERO).minus(((p1).plus(p1)).multipliedBy(p1));
          let _rhs187 = ((_module.__default.fm1((_1151_v99).f10, globalState)) ? (new BigNumber((_1154_v102).length)) : ((p1).multipliedBy((_this).f9)));
          let _lhs136 = globalState;
          let _lhs137 = _1149_v97;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1149_v97).length));
          r1 = _rhs185;
          _lhs136.f5 = _rhs186;
          _lhs137[_lhs138] = _rhs187;
          let _1155_v103;
          let _nw212 = new _module.C3();
          _nw212.__ctor((_dafny.ZERO).minus(p0));
          _1155_v103 = _nw212;
          let _1156_v104;
          let _nw213 = Array((new BigNumber(8)).toNumber());
          _nw213[(_dafny.ZERO).toNumber()] = _1155_v103;
          _nw213[(_dafny.ONE).toNumber()] = _1155_v103;
          _nw213[(new BigNumber(2)).toNumber()] = _1155_v103;
          _nw213[(new BigNumber(3)).toNumber()] = _1155_v103;
          _nw213[(new BigNumber(4)).toNumber()] = _1155_v103;
          _nw213[(new BigNumber(5)).toNumber()] = _1155_v103;
          _nw213[(new BigNumber(6)).toNumber()] = _1155_v103;
          _nw213[(new BigNumber(7)).toNumber()] = _1155_v103;
          _1156_v104 = _nw213;
          let _1157_v105;
          _1157_v105 = _dafny.Set.fromElements(_1076_v46);
          let _index144 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1149_v97).length));
          let _rhs188 = (((_1003_v1) ? ((_this).f9) : (new BigNumber((_1157_v105).length)))).isLessThanOrEqualTo(new BigNumber(-348));
          let _rhs189 = !_dafny.areEqual(_1147_v96, _1147_v96);
          let _rhs190 = _1156_v104;
          let _rhs191 = _module.__default.fm27(globalState);
          let _rhs192 = (_1149_v97)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_1149_v97).length))];
          let _lhs139 = globalState;
          let _lhs140 = _1149_v97;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1149_v97).length));
          _lhs139.f2 = _rhs188;
          _1003_v1 = _rhs189;
          _1156_v104 = _rhs190;
          _1002_v0 = _rhs191;
          _lhs140[_lhs141] = _rhs192;
        }
        let _1158_v106;
        let _init37 = ((_1159_v1) => function (_1160_i13) {
          return _1159_v1;
        })(_1003_v1);
        let _nw214 = Array((_dafny.ONE).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw214.length); _i0_37++) {
          _nw214[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1158_v106 = _nw214;
        let _index145 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_1158_v106).length));
        (_1158_v106)[_index145] = ((_1003_v1) ? (_1003_v1) : (false));
        let _index146 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_1158_v106).length));
        (_1158_v106)[_index146] = _1003_v1;
      } else {
        let _1161_v107;
        _1161_v107 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1162_v108;
        _1162_v108 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1161_v107);
        (globalState).f5 = (new BigNumber((_1162_v108).length)).multipliedBy(new BigNumber(535));
        r1 = (_dafny.ZERO).minus(p0);
        let _1163_v109;
        _1163_v109 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1003_v1);
        let _1164_v110;
        _1164_v110 = _module.D0.create_DC0(_1163_v109);
        let _1165_v111;
        _1165_v111 = _module.D0.create_DC2(_1164_v110);
        let _1166_v112;
        _1166_v112 = _module.D0.create_DC2(_1165_v111);
        let _source17 = _module.D0.create_DC2(((_1003_v1) ? (_1165_v111) : (_1165_v111)));
        if (_source17.is_DC1) {
          let _1167___mcc_h11 = (_source17).cf1;
          let _1168_cf1 = _1167___mcc_h11;
          let _1169_v113;
          _1169_v113 = _dafny.Set.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), _1161_v107);
          let _1170_v114;
          _1170_v114 = _dafny.Seq.of(p0);
          let _1171_v115;
          _1171_v115 = _dafny.Seq.of(_1168_cf1, false);
          let _1172_v116;
          _1172_v116 = _dafny.Set.fromElements(_1003_v1, _1168_cf1);
          let _1173_v117;
          _1173_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1172_v116,_1168_cf1);
          let _1174_v118;
          _1174_v118 = _dafny.Seq.of((_1173_v117).update(_dafny.Set.fromElements(_1003_v1), true), _1173_v117, _dafny.Map.Empty.slice().updateUnsafe(_1172_v116,_1168_cf1));
          let _1175_v119;
          _1175_v119 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(889),p1);
          let _1176_v120;
          _1176_v120 = _dafny.Seq.of(_1002_v0);
          let _1177_v121;
          _1177_v121 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(406)), ((_1178_v107) => function (_1179_i14) {
            return _1178_v107;
          })(_1161_v107)), (_1176_v120)[_module.__default.safeIndex(p0, new BigNumber((_1176_v120).length))], _1002_v0, _dafny.Seq.UnicodeFromString("nikx"), _1002_v0);
          let _1180_v122;
          let _nw215 = Array((new BigNumber(25)).toNumber());
          _nw215[(_dafny.ZERO).toNumber()] = p0;
          _nw215[(_dafny.ONE).toNumber()] = (((((_1163_v109).contains(new BigNumber((_1169_v113).length))) ? ((_1163_v109).get(new BigNumber((_1169_v113).length))) : (_1003_v1))) ? (_module.__default.fm0((_this).f9, (_this).f9, p0, _1168_cf1, globalState)) : ((_this).f9));
          _nw215[(new BigNumber(2)).toNumber()] = (_this).f9;
          _nw215[(new BigNumber(3)).toNumber()] = (_1170_v114)[_module.__default.safeIndex((_this).f9, new BigNumber((_1170_v114).length))];
          _nw215[(new BigNumber(4)).toNumber()] = ((_1003_v1) ? ((_this).f9) : ((_this).f9));
          _nw215[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw215[(new BigNumber(6)).toNumber()] = _module.__default.fm0(new BigNumber((_1171_v115).length), new BigNumber(-482), p1, true, globalState);
          _nw215[(new BigNumber(7)).toNumber()] = (_this).f9;
          _nw215[(new BigNumber(8)).toNumber()] = p1;
          _nw215[(new BigNumber(9)).toNumber()] = p0;
          _nw215[(new BigNumber(10)).toNumber()] = (p0).multipliedBy(new BigNumber((_dafny.Set.fromElements((_this).f9)).length));
          _nw215[(new BigNumber(11)).toNumber()] = p0;
          _nw215[(new BigNumber(12)).toNumber()] = new BigNumber(((_1174_v118)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1170_v114).length)), new BigNumber((_1174_v118).length))]).length);
          _nw215[(new BigNumber(13)).toNumber()] = p0;
          _nw215[(new BigNumber(14)).toNumber()] = new BigNumber(((_1175_v119).Merge(_1175_v119)).length);
          _nw215[(new BigNumber(15)).toNumber()] = (_this).f9;
          _nw215[(new BigNumber(16)).toNumber()] = (new BigNumber((_1171_v115).length)).multipliedBy(new BigNumber((_1172_v116).length));
          _nw215[(new BigNumber(17)).toNumber()] = p1;
          _nw215[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.update((_1177_v121)[_module.__default.safeIndex(_module.__default.fm0(new BigNumber((_1002_v0).length), new BigNumber(378), (_this).f9, true, globalState), new BigNumber((_1177_v121).length))], _module.__default.safeIndex(new BigNumber(571), new BigNumber(((_1177_v121)[_module.__default.safeIndex(_module.__default.fm0(new BigNumber((_1002_v0).length), new BigNumber(378), (_this).f9, true, globalState), new BigNumber((_1177_v121).length))]).length)), _1161_v107)).length);
          _nw215[(new BigNumber(19)).toNumber()] = p1;
          _nw215[(new BigNumber(20)).toNumber()] = p0;
          _nw215[(new BigNumber(21)).toNumber()] = (_this).f9;
          _nw215[(new BigNumber(22)).toNumber()] = new BigNumber(21);
          _nw215[(new BigNumber(23)).toNumber()] = p1;
          _nw215[(new BigNumber(24)).toNumber()] = p1;
          _1180_v122 = _nw215;
          _1180_v122 = _1180_v122;
          (globalState).f8 = ((_1172_v116).Difference(_1172_v116)).Union((_1172_v116).Union(_1172_v116));
          let _1181_v123;
          let _nw216 = new _module.C3();
          _nw216.__ctor(p0);
          _1181_v123 = _nw216;
          let _1182_v124;
          let _nw217 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1182_v124 = _nw217;
          let _index147 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1182_v124).length));
          (_1182_v124)[_index147] = _1002_v0;
          let _1183_v125;
          let _init38 = ((_1184_v1) => function (_1185_i15) {
            return _1184_v1;
          })(_1003_v1);
          let _nw218 = Array((new BigNumber(19)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw218.length); _i0_38++) {
            _nw218[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1183_v125 = _nw218;
          let _index148 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1183_v125).length));
          (_1183_v125)[_index148] = _1003_v1;
          let _1186_v126;
          _1186_v126 = _dafny.Set.fromElements((_1181_v123).f12);
          let _index149 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1182_v124).length));
          let _index150 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1183_v125).length));
          let _rhs193 = _dafny.Seq.update(_dafny.Seq.Concat(_1002_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-283)), ((_1187_v107) => function (_1188_i16) {
            return _1187_v107;
          })(_1161_v107))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1002_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-283)), ((_1189_v107) => function (_1190_i16) {
            return _1189_v107;
          })(_1161_v107)))).length)), _1161_v107);
          let _rhs194 = _1003_v1;
          let _rhs195 = _1003_v1;
          let _rhs196 = new BigNumber(292);
          let _rhs197 = ((_1186_v126).Difference(_1186_v126)).IsSubsetOf(_1186_v126);
          let _lhs142 = _1182_v124;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1182_v124).length));
          let _lhs144 = _1183_v125;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1183_v125).length));
          let _lhs146 = globalState;
          let _lhs147 = globalState;
          _lhs142[_lhs143] = _rhs193;
          r2 = _rhs194;
          _lhs144[_lhs145] = _rhs195;
          _lhs146.f5 = _rhs196;
          _lhs147.f2 = _rhs197;
        } else if (_source17.is_DC0) {
          let _1191___mcc_h12 = (_source17).cf0;
          let _1192_cf0 = _1191___mcc_h12;
          let _1193_v127;
          let _init39 = ((_1194_v1) => function (_1195_i17) {
            return !(_1194_v1) || (false);
          })(_1003_v1);
          let _nw219 = Array((new BigNumber(3)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw219.length); _i0_39++) {
            _nw219[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1193_v127 = _nw219;
          let _1196_v128;
          _1196_v128 = _module.D4.create_DC11(_1193_v127);
          let _pat_let_tv38 = _1193_v127;
          _1193_v127 = (function (_pat_let20_0) {
            return function (_1197_dt__update__tmp_h2) {
              return function (_pat_let21_0) {
                return function (_1198_dt__update_hcf13_h0) {
                  return _module.D4.create_DC11(_1198_dt__update_hcf13_h0);
                }(_pat_let21_0);
              }(_pat_let_tv38);
            }(_pat_let20_0);
          }(_1196_v128)).dtor_cf13;
          let _1199_v129;
          let _nw220 = new _module.C1();
          _nw220.__ctor();
          _1199_v129 = _nw220;
          let _1200_v130;
          let _nw221 = new _module.C3();
          _nw221.__ctor((_dafny.ZERO).minus((_this).f9));
          _1200_v130 = _nw221;
          _1192_cf0 = (_1192_cf0).update((_1200_v130).f12, _1003_v1);
        } else {
          let _1201___mcc_h13 = (_source17).cf2;
          let _1202_cf2 = _1201___mcc_h13;
          (globalState).f5 = (_this).f9;
          (globalState).f5 = p1;
          r1 = (_this).f9;
          let _1203_v131;
          let _nw222 = new _module.C1();
          _nw222.__ctor();
          _1203_v131 = _nw222;
        }
        let _1204_v132;
        let _nw223 = Array((_dafny.ONE).toNumber());
        _nw223[(_dafny.ZERO).toNumber()] = !(true) || (_1003_v1);
        _1204_v132 = _nw223;
        let _1205_v133;
        _1205_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1002_v0,_1003_v1);
        let _1206_v134;
        _1206_v134 = _dafny.MultiSet.fromElements(_1003_v1, _1003_v1);
        let _index151 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1204_v132).length));
        (_1204_v132)[_index151] = (((_1205_v133).contains(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("rngu"), _module.__default.safeIndex(new BigNumber((_1206_v134).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("rngu")).length)), _1161_v107))) ? ((_1205_v133).get(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("rngu"), _module.__default.safeIndex(new BigNumber((_1206_v134).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("rngu")).length)), _1161_v107))) : (_1003_v1));
        let _index152 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1204_v132).length));
        (_1204_v132)[_index152] = _1003_v1;
        r1 = new BigNumber((_1206_v134).cardinality());
      }
      let _1207_v135;
      let _nw224 = new _module.C1();
      _nw224.__ctor();
      _1207_v135 = _nw224;
      let _1208_v136;
      _1208_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v135,_1003_v1);
      let _1209_v137;
      _1209_v137 = _dafny.Map.Empty.slice().updateUnsafe(_1003_v1,_1003_v1);
      let _1210_v138;
      _1210_v138 = new _dafny.CodePoint('o'.codePointAt(0));
      let _1211_v139;
      _1211_v139 = _module.D10.create_DC27(_1003_v1, p1, _1210_v138, _dafny.MultiSet.fromElements(_1003_v1));
      let _1212_v140;
      _1212_v140 = _dafny.Seq.of(_1002_v0, _1002_v0, _1002_v0);
      let _1213_v141;
      _1213_v141 = _dafny.Set.fromElements((_1212_v140)[_module.__default.safeIndex(p0, new BigNumber((_1212_v140).length))], _1002_v0);
      let _1214_v142;
      let _nw225 = Array((new BigNumber(27)).toNumber());
      _nw225[(_dafny.ZERO).toNumber()] = !(_1003_v1);
      _nw225[(_dafny.ONE).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(2)).toNumber()] = ((!(!(_1003_v1))) ? ((((_1208_v136).contains(_1207_v135)) ? ((_1208_v136).get(_1207_v135)) : (_1003_v1))) : (false));
      _nw225[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_1003_v1, globalState);
      _nw225[(new BigNumber(4)).toNumber()] = !(_1003_v1) || (_1003_v1);
      _nw225[(new BigNumber(5)).toNumber()] = (_1003_v1) === (_1003_v1);
      _nw225[(new BigNumber(6)).toNumber()] = !(true);
      _nw225[(new BigNumber(7)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(p1,_1003_v1)).equals(_dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm1(_1003_v1, globalState)));
      _nw225[(new BigNumber(8)).toNumber()] = !(_1003_v1) || ((((_1209_v137).contains(_1003_v1)) ? ((_1209_v137).get(_1003_v1)) : (_1003_v1)));
      _nw225[(new BigNumber(9)).toNumber()] = (_module.D9.create_DC24(_1003_v1, _1003_v1, _1003_v1)).dtor_cf35;
      _nw225[(new BigNumber(10)).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(11)).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(12)).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(13)).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(14)).toNumber()] = !(false);
      _nw225[(new BigNumber(15)).toNumber()] = !(false);
      _nw225[(new BigNumber(16)).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(17)).toNumber()] = ((_this).f9).isLessThan((_this).f9);
      _nw225[(new BigNumber(18)).toNumber()] = (_1207_v135).fm4(_1004_v2, _1004_v2, (_1211_v139).dtor_cf41, globalState);
      _nw225[(new BigNumber(19)).toNumber()] = false;
      _nw225[(new BigNumber(20)).toNumber()] = (_1213_v141).IsSubsetOf(_dafny.Set.fromElements(_1002_v0, _1002_v0, _1002_v0, _dafny.Seq.UnicodeFromString("q")));
      _nw225[(new BigNumber(21)).toNumber()] = (p1).isEqualTo((_this).f9);
      _nw225[(new BigNumber(22)).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(23)).toNumber()] = !(_1003_v1) || (_1003_v1);
      _nw225[(new BigNumber(24)).toNumber()] = (_1003_v1) && (_1003_v1);
      _nw225[(new BigNumber(25)).toNumber()] = _1003_v1;
      _nw225[(new BigNumber(26)).toNumber()] = _1003_v1;
      _1214_v142 = _nw225;
      let _index153 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length));
      (_1214_v142)[_index153] = (_1003_v1) && (_1003_v1);
      let _1215_v143;
      let _init40 = ((_1216_v1) => function (_1217_i18) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(912)), function (_1218_i19) {
          return new BigNumber(607);
        })).length),_dafny.Set.fromElements(_1216_v1, _1216_v1))).length), new BigNumber(241));
      })(_1003_v1);
      let _nw226 = Array((new BigNumber(2)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw226.length); _i0_40++) {
        _nw226[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1215_v143 = _nw226;
      let _index154 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1215_v143).length));
      (_1215_v143)[_index154] = _dafny.MultiSet.fromElements(p1);
      let _1219_v144;
      _1219_v144 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_1002_v0, _1002_v0, _1002_v0)).cardinality()), (_this).f9, (_this).f9);
      let _index155 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length));
      let _index156 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1215_v143).length));
      let _rhs198 = _1003_v1;
      let _rhs199 = _1219_v144;
      let _lhs148 = _1214_v142;
      let _lhs149 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length));
      let _lhs150 = _1215_v143;
      let _lhs151 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1215_v143).length));
      _lhs148[_lhs149] = _rhs198;
      _lhs150[_lhs151] = _rhs199;
      let _1220_v145;
      let _nw227 = new _module.C2();
      _nw227.__ctor(_1003_v1);
      _1220_v145 = _nw227;
      let _1221_v146;
      _1221_v146 = _module.D7.create_DC18();
      _1221_v146 = (((_1214_v142)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length))]) ? (_1221_v146) : (_1221_v146));
      let _1222_v147;
      _1222_v147 = _module.D11.create_DC31(p0, (_dafny.ZERO).minus((_this).f9));
      let _index157 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length));
      (_1214_v142)[_index157] = ((_1222_v147).dtor_cf45).isLessThan(p1);
      let _1223_v148;
      _1223_v148 = _dafny.Seq.of((_1220_v145).f13);
      let _1224_v149;
      _1224_v149 = _dafny.Seq.of(_1223_v148, _1223_v148);
      let _1225_v150;
      _1225_v150 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm32(_1224_v149, globalState),true);
      let _1226_v151;
      _1226_v151 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1225_v150).length));
      r0 = _1226_v151;
      let _1227_v152;
      _1227_v152 = _dafny.MultiSet.fromElements((_1214_v142)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length))]);
      let _1228_v153;
      _1228_v153 = _module.D7.create_DC19(_module.D7.create_DC18());
      let _1229_v154;
      _1229_v154 = _dafny.Map.Empty.slice().updateUnsafe(_1228_v153,(_1214_v142)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length))]);
      let _1230_v155;
      _1230_v155 = _dafny.Map.Empty.slice().updateUnsafe(_1212_v140,(_this).f9);
      r1 = (_module.__default.safeDivisionInt(p0, (((_1227_v152).contains((((_1229_v154).contains(_1228_v153)) ? ((_1229_v154).get(_1228_v153)) : ((_1214_v142)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length))])))) ? ((_1227_v152).get((((_1229_v154).contains(_1228_v153)) ? ((_1229_v154).get(_1228_v153)) : ((_1214_v142)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1214_v142).length))])))) : (new BigNumber(-660))))).plus((((_1230_v155).contains(_dafny.Seq.of(_1002_v0))) ? ((_1230_v155).get(_dafny.Seq.of(_1002_v0))) : (p0)));
      let _1231_v156;
      _1231_v156 = _dafny.Set.fromElements(_1214_v142);
      r2 = (_module.__default.safeModuloInt(new BigNumber((_1231_v156).length), new BigNumber(326))).isEqualTo(p1);
      return [r0, r1, r2];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
