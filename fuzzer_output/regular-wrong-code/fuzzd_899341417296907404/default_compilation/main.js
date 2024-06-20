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
    static fm0(p0, p1, p2, globalState) {
      return (new BigNumber((_dafny.Seq.of(!(!(!(false))), true, true, false, false)).length)).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tialul"), _dafny.Seq.UnicodeFromString("olu"))).length));
    };
    static fm1(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()),true)).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(552), new BigNumber(799))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(552)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(799)))) {
            _coll0.push([_module.__default.safeDivisionInt(_0_v0, new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("vi"))).cardinality())),(_module.D0.create_DC1(_dafny.Seq.UnicodeFromString("nqxdlqnf"), (_module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_1_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
}), true, true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).dtor_cf3, false, new BigNumber(326))).dtor_cf3]);
          }
        }
        return _coll0;
      }())).contains(_module.__default.safeModuloInt(new BigNumber(859), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(246),new BigNumber(396))).length)));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('b'.codePointAt(0));
    };
    static fm5(p0, p1, p2, globalState) {
      let _source0 = _module.D0.create_DC2(_dafny.Seq.Create(_module.__default.abs(new BigNumber(118)), function (_2_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}), !(false), _dafny.Seq.of(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(948),true)).length), new _dafny.CodePoint('v'.codePointAt(0)));
      if (_source0.is_DC1) {
        let _3___mcc_h0 = (_source0).cf1;
        let _4___mcc_h1 = (_source0).cf2;
        let _5___mcc_h2 = (_source0).cf3;
        let _6___mcc_h3 = (_source0).cf4;
        let _7_cf4 = _6___mcc_h3;
        let _8_cf3 = _5___mcc_h2;
        let _9_cf2 = _4___mcc_h1;
        let _10_cf1 = _3___mcc_h0;
        if (_8_cf3) {
          return _dafny.Seq.of(_7_cf4);
        } else {
          return _dafny.Seq.of(_7_cf4, _7_cf4);
        }
      } else if (_source0.is_DC2) {
        let _11___mcc_h4 = (_source0).cf5;
        let _12___mcc_h5 = (_source0).cf6;
        let _13___mcc_h6 = (_source0).cf7;
        let _14___mcc_h7 = (_source0).cf8;
        let _15___mcc_h8 = (_source0).cf9;
        let _16_cf9 = _15___mcc_h8;
        let _17_cf8 = _14___mcc_h7;
        let _18_cf7 = _13___mcc_h6;
        let _19_cf6 = _12___mcc_h5;
        let _20_cf5 = _11___mcc_h4;
        return _dafny.Seq.of(_17_cf8);
      } else if (_source0.is_DC3) {
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(855),!(false))).length), new BigNumber(794), new BigNumber(211), new BigNumber(165), new BigNumber((_dafny.Seq.UnicodeFromString("hrwsevras")).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("fqrxt")).length)));
      } else if (_source0.is_DC0) {
        let _21___mcc_h9 = (_source0).cf0;
        let _22_cf0 = _21___mcc_h9;
        return _dafny.Seq.of(new BigNumber(734));
      } else {
        let _23___mcc_h10 = (_source0).cf10;
        let _24_cf10 = _23___mcc_h10;
        return _dafny.Seq.of(new BigNumber(911), new BigNumber(62));
      }
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), function (_25_i0) {
  return new _dafny.CodePoint('k'.codePointAt(0));
}), !(false) || (true), !((_dafny.MultiSet.fromElements(true)).IsSubsetOf(_dafny.MultiSet.fromElements(true, false))), new BigNumber(895));
    };
    static fm7(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(438)), function (_26_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-401))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(138),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(889)), function (_27_i1) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length))).length))).length), new BigNumber(556))).cardinality()))));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), function (_28_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length))).Union(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-292))).length))))).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("my")).length)));
    };
    static fm9(p0, p1, p2, globalState) {
      return (((false) ? (function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("s"))).Elements) {
          let _29_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("s")), _29_v0)) {
            _coll1.add(_29_v0);
          }
        }
        return _coll1;
      }()) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ncxyqxudm"))))).Union(((false) ? (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("pyfiuigns"))) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("wruwanx"), _dafny.Seq.UnicodeFromString("mpqo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(137)), function (_30_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })))));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _31_v0;
      _31_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _32_v1;
      _32_v1 = _dafny.MultiSet.fromElements(!(_31_v0).equals(_31_v0), !(p0));
      (globalState).f4 = (((_32_v1).contains(p0)) ? ((_32_v1).get(p0)) : (_module.__default.fm0((_dafny.ZERO).minus(p3), p3, p0, globalState)));
      let _33_v2;
      let _nw0 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _33_v2 = _nw0;
      let _34_v3;
      _34_v3 = new _dafny.CodePoint('s'.codePointAt(0));
      let _index0 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_33_v2).length));
      (_33_v2)[_index0] = _34_v3;
      let _35_v4;
      _35_v4 = _dafny.MultiSet.fromElements(new BigNumber(163), p1);
      let _36_v5;
      let _nw1 = new _module.C0();
      _nw1.__ctor();
      _36_v5 = _nw1;
      let _37_v6;
      _37_v6 = _module.D4.create_DC12(_34_v3, _35_v4, p0, p1, _36_v5);
      let _index1 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_33_v2).length));
      (_33_v2)[_index1] = (_37_v6).dtor_cf16;
      let _38_v7;
      let _init0 = function (_39_i0) {
        return _module.__default.safeDivisionInt(_39_i0, new BigNumber(824));
      };
      let _nw2 = Array((new BigNumber(16)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _38_v7 = _nw2;
      _38_v7 = _38_v7;
      let _40_v8;
      _40_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _41_v9;
      _41_v9 = _dafny.Seq.of(p0);
      let _42_v10;
      _42_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_41_v9);
      let _43_v11;
      _43_v11 = _42_v10;
      let _pat_let_tv0 = _40_v8;
      _40_v8 = function (_source1) {
        let _44___mcc_h0 = _source1;
        let _45_cf13 = _44___mcc_h0;
        return _pat_let_tv0;
      }(_43_v11);
      let _46_v12;
      _46_v12 = _dafny.Seq.of((_33_v2)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_33_v2).length))]);
      let _47_v13;
      _47_v13 = _dafny.Seq.of(_46_v12, _46_v12);
      let _hi0 = new BigNumber((_47_v13).length);
      for (let _48_i1 = new BigNumber((_41_v9).length); _48_i1.isLessThan(_hi0); _48_i1 = _48_i1.plus(_dafny.ONE)) {
        (globalState).f4 = (_module.__default.fm0(p3, new BigNumber((_41_v9).length), !(false), globalState)).multipliedBy(_48_i1);
        let _49_v14;
        _49_v14 = _dafny.Map.Empty.slice().updateUnsafe(_46_v12,p0);
        let _50_v15;
        _50_v15 = _dafny.Set.fromElements(p1, p3);
        (globalState).f1 = (((_49_v14).contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(624)), ((_53_v2) => function (_54_i2) {
          return (_53_v2)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_53_v2).length))];
        })(_33_v2)), _46_v12))) ? ((_49_v14).get(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(624)), ((_51_v2) => function (_52_i2) {
          return (_51_v2)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_51_v2).length))];
        })(_33_v2)), _46_v12))) : (!(!(_50_v15).contains(_48_i1))));
        (globalState).f4 = _module.__default.fm0((new BigNumber(501)).multipliedBy(p3), new BigNumber(-702), p0, globalState);
        (globalState).f5 = !(p0);
      }
      let _hi1 = new BigNumber((_dafny.Seq.UnicodeFromString("hnssgkhut")).length);
      for (let _55_i3 = p3; _55_i3.isLessThan(_hi1); _55_i3 = _55_i3.plus(_dafny.ONE)) {
        (globalState).f4 = p3;
        if (p0) {
          (globalState).f5 = p0;
          let _56_v16;
          _56_v16 = _module.D1.create_DC6(p1);
          _56_v16 = _56_v16;
          let _57_v17;
          _57_v17 = _module.D0.create_DC0(new BigNumber(843));
          let _58_v18;
          _58_v18 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(545)), ((_59_p3) => function (_60_i4) {
            return _59_p3;
          })(p3)));
          let _61_v19;
          _61_v19 = _dafny.Seq.of(_57_v17, _module.__default.fm6(new BigNumber(((_58_v18)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_58_v18).length))]).length), _40_v8, p0, p0, globalState));
          let _62_v20;
          let _nw3 = Array((new BigNumber(11)).toNumber());
          _nw3[(_dafny.ZERO).toNumber()] = _61_v19;
          _nw3[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_61_v19, _61_v19);
          _nw3[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), ((_63_v17) => function (_64_i5) {
            return _63_v17;
          })(_57_v17));
          _nw3[(new BigNumber(3)).toNumber()] = _61_v19;
          _nw3[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_61_v19, _61_v19);
          _nw3[(new BigNumber(5)).toNumber()] = _61_v19;
          _nw3[(new BigNumber(6)).toNumber()] = _61_v19;
          _nw3[(new BigNumber(7)).toNumber()] = _61_v19;
          _nw3[(new BigNumber(8)).toNumber()] = _61_v19;
          _nw3[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(965)), ((_65_v17) => function (_66_i6) {
            return _65_v17;
          })(_57_v17));
          _nw3[(new BigNumber(10)).toNumber()] = _61_v19;
          _62_v20 = _nw3;
          let _67_v21;
          let _nw4 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
          _67_v21 = _nw4;
          let _68_v22;
          _68_v22 = _67_v21;
          let _rhs0 = _62_v20;
          let _rhs1 = (_68_v22);
          let _rhs2 = _38_v7;
          _62_v20 = _rhs0;
          _67_v21 = _rhs1;
          _38_v7 = _rhs2;
          let _index2 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_38_v7).length));
          (_38_v7)[_index2] = (_dafny.ZERO).minus(p3);
          let _index3 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_38_v7).length));
          (_38_v7)[_index3] = _module.__default.safeDivisionInt(p1, _55_i3);
          let _69_v23;
          _69_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
          let _70_v24;
          _70_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(218),p3);
          let _71_v25;
          _71_v25 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_72_p3) => function (_73_i7) {
            return _72_p3;
          })(p3))).length));
          let _74_v26;
          _74_v26 = _dafny.Seq.of(_36_v5, _36_v5);
          let _index4 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_38_v7).length));
          let _index5 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_38_v7).length));
          let _rhs3 = (((_69_v23).contains(!(_module.__default.fm0((((_70_v24).contains(p3)) ? ((_70_v24).get(p3)) : (_55_i3)), p3, p0, globalState)).isEqualTo(p3))) ? ((_69_v23).get(!(_module.__default.fm0((((_70_v24).contains(p3)) ? ((_70_v24).get(p3)) : (_55_i3)), p3, p0, globalState)).isEqualTo(p3))) : (p1));
          let _rhs4 = _module.__default.safeModuloInt(p3, _module.__default.fm0((p2)[_module.__default.safeIndex(new BigNumber(474), new BigNumber((p2).length))], _55_i3, (_37_v6).dtor_cf18, globalState));
          let _rhs5 = !_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Concat(_41_v9, _41_v9), _module.__default.safeIndex((_dafny.ZERO).minus(_55_i3), new BigNumber((_dafny.Seq.Concat(_41_v9, _41_v9)).length)), _module.__default.fm1(p2, _71_v25, p0, globalState)), (_32_v1).IsProperSubsetOf(_32_v1));
          let _rhs6 = p3;
          let _rhs7 = (_dafny.ZERO).minus((_55_i3).plus(_module.__default.safeDivisionInt(p1, new BigNumber((_74_v26).length))));
          let _lhs0 = globalState;
          let _lhs1 = _38_v7;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_38_v7).length));
          let _lhs3 = globalState;
          let _lhs4 = _38_v7;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_38_v7).length));
          let _lhs6 = globalState;
          _lhs0.f4 = _rhs3;
          _lhs1[_lhs2] = _rhs4;
          _lhs3.f1 = _rhs5;
          _lhs4[_lhs5] = _rhs6;
          _lhs6.f4 = _rhs7;
          (globalState).f4 = (_module.D4.create_DC12((_33_v2)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_33_v2).length))], _35_v4, !(false), new BigNumber((_40_v8).length), _36_v5)).dtor_cf19;
        } else {
          (globalState).f5 = (_41_v9)[_module.__default.safeIndex(p3, new BigNumber((_41_v9).length))];
          let _rhs8 = p0;
          let _rhs9 = p0;
          let _rhs10 = _module.__default.safeDivisionInt(new BigNumber(939), (p1).minus(new BigNumber((_46_v12).length)));
          let _lhs7 = globalState;
          let _lhs8 = globalState;
          let _lhs9 = globalState;
          _lhs7.f1 = _rhs8;
          _lhs8.f5 = _rhs9;
          _lhs9.f4 = _rhs10;
          let _75_v27;
          let _init1 = ((_76_v9, _77_p3, _78_v2, _79_p0) => function (_80_i8) {
            return (_dafny.Map.Empty.slice().updateUnsafe(!((_76_v9)[_module.__default.safeIndex(_77_p3, new BigNumber((_76_v9).length))]),(_78_v2)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_78_v2).length))])).Merge(_dafny.Map.Empty.slice().updateUnsafe(_79_p0,new _dafny.CodePoint('e'.codePointAt(0))));
          })(_41_v9, p3, _33_v2, p0);
          let _nw5 = Array((new BigNumber(26)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
            _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _75_v27 = _nw5;
          let _81_v28;
          _81_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('c'.codePointAt(0)));
          let _index6 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_75_v27).length));
          (_75_v27)[_index6] = _81_v28;
          let _82_v29;
          _82_v29 = _dafny.Map.Empty.slice().updateUnsafe((_40_v8).Merge(_40_v8),_81_v28);
          let _index7 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_75_v27).length));
          let _rhs11 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(_55_i3, (((_35_v4).contains(p3)) ? ((_35_v4).get(p3)) : (new BigNumber((_41_v9).length))), p0, globalState)));
          let _rhs12 = (((_82_v29).contains(_40_v8)) ? ((_82_v29).get(_40_v8)) : ((_81_v28).update(p0, _34_v3)));
          let _lhs10 = globalState;
          let _lhs11 = _75_v27;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_75_v27).length));
          _lhs10.f4 = _rhs11;
          _lhs11[_lhs12] = _rhs12;
          _36_v5 = _36_v5;
          (globalState).f5 = p0;
        }
        (globalState).f5 = p0;
        let _83_v30;
        let _nw6 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _83_v30 = _nw6;
        let _index8 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_83_v30).length));
        (_83_v30)[_index8] = _dafny.Seq.UnicodeFromString("ouixcdf");
        let _index9 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_83_v30).length));
        let _rhs13 = _dafny.Seq.Concat(_46_v12, _46_v12);
        let _rhs14 = ((p1).multipliedBy(_55_i3)).isEqualTo(_55_i3);
        let _lhs13 = _83_v30;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_83_v30).length));
        let _lhs15 = globalState;
        _lhs13[_lhs14] = _rhs13;
        _lhs15.f5 = _rhs14;
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _84_v0;
      _84_v0 = new BigNumber(624);
      let _85_v1;
      _85_v1 = _dafny.Seq.UnicodeFromString("vo");
      let _86_v2;
      let _nw7 = Array((new BigNumber(12)).toNumber());
      _nw7[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_84_v0);
      _nw7[(_dafny.ONE).toNumber()] = _84_v0;
      _nw7[(new BigNumber(2)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(3)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(4)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(5)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(6)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(7)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(8)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(9)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(10)).toNumber()] = _84_v0;
      _nw7[(new BigNumber(11)).toNumber()] = new BigNumber((_85_v1).length);
      _86_v2 = _nw7;
      let _87_globalState;
      let _nw8 = new _module.GlobalState();
      _nw8.__ctor(true, false, _86_v2, new BigNumber(965), new BigNumber(-64), false);
      _87_globalState = _nw8;
      (_87_globalState).f4 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_84_v0, _84_v0));
      let _index10 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
      (_86_v2)[_index10] = _84_v0;
      let _index11 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
      (_86_v2)[_index11] = _84_v0;
      let _hi2 = new BigNumber((_85_v1).length);
      for (let _88_i0 = new BigNumber((_85_v1).length); _88_i0.isLessThan(_hi2); _88_i0 = _88_i0.plus(_dafny.ONE)) {
        let _89_v3;
        _89_v3 = false;
        let _90_v4;
        let _nw9 = Array((new BigNumber(17)).toNumber());
        _nw9[(_dafny.ZERO).toNumber()] = _89_v3;
        _nw9[(_dafny.ONE).toNumber()] = _89_v3;
        _nw9[(new BigNumber(2)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(3)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(4)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(5)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(6)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(7)).toNumber()] = false;
        _nw9[(new BigNumber(8)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(9)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(10)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(11)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(12)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(13)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(14)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(15)).toNumber()] = _89_v3;
        _nw9[(new BigNumber(16)).toNumber()] = _89_v3;
        _90_v4 = _nw9;
        let _91_v5;
        _91_v5 = _dafny.Map.Empty.slice().updateUnsafe(_88_i0,_90_v4);
        let _92_v6;
        _92_v6 = _dafny.Map.Empty.slice().updateUnsafe(_88_i0,_module.__default.fm0(_88_i0, _84_v0, _89_v3, _87_globalState));
        let _93_v7;
        _93_v7 = _dafny.Seq.of(new BigNumber((_92_v6).length));
        _module.__default.m0(_89_v3, (new BigNumber((_91_v5).length)).minus(_84_v0), _93_v7, ((_89_v3) ? ((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]) : (_88_i0)), _87_globalState);
        let _94_v8;
        _94_v8 = _dafny.Set.fromElements((_dafny.ZERO).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]));
        let _95_v9;
        _95_v9 = _dafny.Seq.of(_89_v3, _module.__default.fm1(_93_v7, _94_v8, _89_v3, _87_globalState), _89_v3, _89_v3);
        let _96_v10;
        _96_v10 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_93_v7, _94_v8, _89_v3, _87_globalState),_dafny.Seq.Concat(_95_v9, _95_v9));
        _96_v10 = (_96_v10).update(_89_v3, _95_v9);
        let _97_v11;
        _97_v11 = _dafny.Map.Empty.slice().updateUnsafe(_89_v3,_89_v3);
        let _98_v12;
        _98_v12 = _dafny.MultiSet.fromElements(_88_i0, _88_i0, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], _module.__default.safeModuloInt(_88_i0, new BigNumber((_97_v11).length)));
        (_87_globalState).f4 = (((_98_v12).contains(new BigNumber((_dafny.Seq.of(_86_v2)).length))) ? ((_98_v12).get(new BigNumber((_dafny.Seq.of(_86_v2)).length))) : (new BigNumber(568)));
        let _index12 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        (_86_v2)[_index12] = new BigNumber(955);
      }
      let _99_v13;
      let _nw10 = new _module.C0();
      _nw10.__ctor();
      _99_v13 = _nw10;
      let _100_v14;
      let _nw11 = Array((new BigNumber(7)).toNumber());
      _100_v14 = _nw11;
      let _index13 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length));
      (_100_v14)[_index13] = _99_v13;
      let _index14 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length));
      (_100_v14)[_index14] = _99_v13;
      let _101_v15;
      _101_v15 = _module.D0.create_DC1(_85_v1, !(false), false, _84_v0);
      let _index15 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
      let _index16 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
      let _rhs15 = _86_v2;
      let _rhs16 = (_101_v15).dtor_cf4;
      let _rhs17 = (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))];
      let _rhs18 = _84_v0;
      let _lhs16 = _86_v2;
      let _lhs17 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
      let _lhs18 = _86_v2;
      let _lhs19 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
      _86_v2 = _rhs15;
      _84_v0 = _rhs16;
      _lhs16[_lhs17] = _rhs17;
      _lhs18[_lhs19] = _rhs18;
      let _102_i1;
      _102_i1 = _dafny.ZERO;
      L0: {
        while (((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]).isLessThan(((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]).plus(_84_v0))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_102_i1)) {
              break L0;
            }
            _102_i1 = (_102_i1).plus(_dafny.ONE);
            let _103_v16;
            _103_v16 = _module.D0.create_DC0((_dafny.ZERO).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]));
            _103_v16 = _103_v16;
            let _104_v17;
            let _nw12 = new _module.C0();
            _nw12.__ctor();
            _104_v17 = _nw12;
            (_87_globalState).f4 = (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))];
            let _105_v18;
            _105_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(775),_84_v0);
            let _106_v19;
            _106_v19 = _dafny.Set.fromElements((_dafny.ZERO).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]));
            (_87_globalState).f5 = (_106_v19).contains((((_105_v18).contains(_84_v0)) ? ((_105_v18).get(_84_v0)) : (_84_v0)));
          }
        }
      }
      let _107_v20;
      _107_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_108_i2) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }),false);
      let _109_v21;
      _109_v21 = _dafny.MultiSet.fromElements(_84_v0);
      let _110_v22;
      _110_v22 = _dafny.Set.fromElements(new BigNumber(-744), new BigNumber((_85_v1).length), new BigNumber((_109_v21).cardinality()), (_dafny.ZERO).minus((((_109_v21).contains((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])) ? ((_109_v21).get((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])) : ((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]))), _84_v0);
      let _111_v23;
      _111_v23 = false;
      if (_module.__default.fm1(_dafny.Seq.of(_84_v0, _84_v0, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], new BigNumber((_107_v20).length), (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]), _110_v22, _111_v23, _87_globalState)) {
        let _112_v24;
        let _nw13 = Array((_dafny.ONE).toNumber()).fill(false);
        _112_v24 = _nw13;
        _112_v24 = _112_v24;
        let _113_v25;
        _113_v25 = _dafny.Seq.of(new BigNumber(-239));
        if ((_111_v23) || (_module.__default.fm1(_113_v25, _dafny.Set.fromElements(new BigNumber(582)), _111_v23, _87_globalState))) {
          let _114_v26;
          _114_v26 = _dafny.Seq.of(_111_v23);
          let _index17 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
          (_86_v2)[_index17] = new BigNumber((_114_v26).length);
          let _115_v27;
          _115_v27 = _dafny.Map.Empty.slice().updateUnsafe(_111_v23,(_110_v22).IsProperSubsetOf(_110_v22));
          _115_v27 = (_115_v27).update(_111_v23, !(_111_v23));
          let _116_v28;
          let _nw14 = Array((new BigNumber(3)).toNumber());
          _nw14[(_dafny.ZERO).toNumber()] = _85_v1;
          _nw14[(_dafny.ONE).toNumber()] = _85_v1;
          _nw14[(new BigNumber(2)).toNumber()] = _85_v1;
          _116_v28 = _nw14;
          let _index18 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_116_v28).length));
          (_116_v28)[_index18] = _85_v1;
          let _index19 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_116_v28).length));
          (_116_v28)[_index19] = _85_v1;
          (_87_globalState).f4 = _84_v0;
          let _117_v29;
          _117_v29 = new _dafny.CodePoint('g'.codePointAt(0));
          let _118_v30;
          _118_v30 = _dafny.Map.Empty.slice().updateUnsafe(_85_v1,_117_v29);
          _118_v30 = (_118_v30).update(_dafny.Seq.Concat((_116_v28)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_116_v28).length))], _85_v1), ((_111_v23) ? (_module.__default.fm4(!(_111_v23), _84_v0, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], _87_globalState)) : (new _dafny.CodePoint('m'.codePointAt(0)))));
        } else {
          let _119_v31;
          _119_v31 = _dafny.Seq.of((_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))], _99_v13);
          (_87_globalState).f4 = (_dafny.ZERO).minus((_113_v25)[_module.__default.safeIndex(new BigNumber((_119_v31).length), new BigNumber((_113_v25).length))]);
          _111_v23 = ((_111_v23) ? (((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]).isLessThan(_84_v0)) : (_module.__default.fm1(_113_v25, _dafny.Set.fromElements(_84_v0), false, _87_globalState)));
          let _120_v32;
          _120_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_110_v22).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_101_v15,(_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])).length)),(_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
          _120_v32 = _120_v32;
          let _121_v33;
          _121_v33 = _dafny.Seq.of(_111_v23);
          let _122_v34;
          _122_v34 = new _dafny.CodePoint('b'.codePointAt(0));
          let _123_v35;
          _123_v35 = _module.D0.create_DC2(_dafny.Seq.UnicodeFromString("cdwg"), _111_v23, _121_v33, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], _122_v34);
          (_87_globalState).f4 = (_123_v35).dtor_cf8;
          (_87_globalState).f1 = (_111_v23) === (_111_v23);
        }
        let _index20 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_112_v24).length));
        (_112_v24)[_index20] = true;
        let _index21 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_112_v24).length));
        (_112_v24)[_index21] = !(_111_v23);
        let _124_v36;
        _124_v36 = _module.D0.create_DC0(_84_v0);
        let _125_v37;
        _125_v37 = _dafny.Map.Empty.slice().updateUnsafe(_111_v23,_111_v23);
        let _126_v38;
        _126_v38 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_113_v25).length)).plus(_84_v0),(_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
        let _index22 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_112_v24).length));
        let _index23 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_112_v24).length));
        let _index24 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        let _rhs19 = _dafny.Seq.IsPrefixOf(_113_v25, _module.__default.fm5((_124_v36).dtor_cf0, _111_v23, (((_125_v37).contains(_111_v23)) ? ((_125_v37).get(_111_v23)) : (_111_v23)), _87_globalState));
        let _rhs20 = _111_v23;
        let _rhs21 = (_84_v0).plus((_84_v0).plus(new BigNumber(634)));
        let _rhs22 = (((_126_v38).contains((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])) ? ((_126_v38).get((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])) : (_84_v0));
        let _rhs23 = true;
        let _lhs20 = _112_v24;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_112_v24).length));
        let _lhs22 = _112_v24;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_112_v24).length));
        let _lhs24 = _86_v2;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        _lhs20[_lhs21] = _rhs19;
        _lhs22[_lhs23] = _rhs20;
        _84_v0 = _rhs21;
        _lhs24[_lhs25] = _rhs22;
        _111_v23 = _rhs23;
        let _127_v39;
        _127_v39 = _dafny.MultiSet.fromElements((_112_v24)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_112_v24).length))], _111_v23, _111_v23);
        let _index25 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        (_86_v2)[_index25] = _module.__default.safeModuloInt(new BigNumber(111), (((_127_v39).contains(_111_v23)) ? ((_127_v39).get(_111_v23)) : ((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])));
        let _128_v40;
        _128_v40 = _dafny.MultiSet.fromElements(_112_v24, _112_v24);
        let _129_v41;
        _129_v41 = _dafny.Map.Empty.slice().updateUnsafe((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))],false);
        let _130_v42;
        _130_v42 = _dafny.Map.Empty.slice().updateUnsafe((_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))],_111_v23);
        let _131_v43;
        _131_v43 = _module.D0.create_DC4(_module.__default.fm6(new BigNumber((_128_v40).cardinality()), _129_v41, (((_130_v42).contains((_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))])) ? ((_130_v42).get((_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))])) : ((_112_v24)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_112_v24).length))])), false, _87_globalState));
        let _132_v44;
        _132_v44 = _dafny.Seq.of(!(true), (_112_v24)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_112_v24).length))]);
        let _133_v45;
        _133_v45 = new _dafny.CodePoint('a'.codePointAt(0));
        let _134_v46;
        _134_v46 = _module.D0.create_DC2(_85_v1, _111_v23, _132_v44, new BigNumber((_dafny.Seq.UnicodeFromString("qra")).length), _133_v45);
        let _pat_let_tv1 = _134_v46;
        _131_v43 = function (_pat_let0_0) {
          return function (_135_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_136_dt__update_hcf10_h0) {
                return _module.D0.create_DC4(_136_dt__update_hcf10_h0);
              }(_pat_let1_0);
            }(_pat_let_tv1);
          }(_pat_let0_0);
        }(_131_v43);
      } else {
        let _index26 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length));
        (_100_v14)[_index26] = _99_v13;
        let _137_v47;
        let _nw15 = new _module.C0();
        _nw15.__ctor();
        _137_v47 = _nw15;
        let _138_v48;
        _138_v48 = _dafny.Seq.of(_84_v0, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], new BigNumber((_85_v1).length));
        let _139_v50;
        _139_v50 = _dafny.Seq.of(_111_v23);
        (_87_globalState).f1 = _module.__default.fm1(_138_v48, (_110_v22).Difference(function () {
          let _coll2 = new _dafny.Set();
          for (const _compr_2 of (_dafny.MultiSet.FromArray(_138_v48)).Elements) {
            let _140_v49 = _compr_2;
            if ((_dafny.MultiSet.FromArray(_138_v48)).contains(_140_v49)) {
              _coll2.add((_140_v49).plus(new BigNumber(426)));
            }
          }
          return _coll2;
        }()), (_111_v23) || ((_139_v50)[_module.__default.safeIndex(_84_v0, new BigNumber((_139_v50).length))]), _87_globalState);
        let _index27 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        (_86_v2)[_index27] = _84_v0;
        let _141_v51;
        _141_v51 = _dafny.Map.Empty.slice().updateUnsafe(_84_v0,_111_v23);
        let _142_v52;
        _142_v52 = new _dafny.CodePoint('s'.codePointAt(0));
        let _143_v53;
        _143_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(281),_142_v52);
        let _144_v54;
        let _nw16 = Array((new BigNumber(9)).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = !(false);
        _nw16[(_dafny.ONE).toNumber()] = !(_module.__default.fm1(_dafny.Seq.of(_84_v0), _110_v22, _111_v23, _87_globalState)) || (_111_v23);
        _nw16[(new BigNumber(2)).toNumber()] = !((_139_v50)[_module.__default.safeIndex(_84_v0, new BigNumber((_139_v50).length))]) || (!(_111_v23));
        _nw16[(new BigNumber(3)).toNumber()] = _111_v23;
        _nw16[(new BigNumber(4)).toNumber()] = (((((_141_v51).contains((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])) ? ((_141_v51).get((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])) : (_111_v23))) ? (_111_v23) : (!(_111_v23)));
        _nw16[(new BigNumber(5)).toNumber()] = true;
        _nw16[(new BigNumber(6)).toNumber()] = _dafny.Seq.IsPrefixOf(_138_v48, _138_v48);
        _nw16[(new BigNumber(7)).toNumber()] = !(_143_v53).equals(_dafny.Map.Empty.slice().updateUnsafe(_84_v0,_142_v52));
        _nw16[(new BigNumber(8)).toNumber()] = _111_v23;
        _144_v54 = _nw16;
        _144_v54 = _144_v54;
      }
      let _145_v55;
      _145_v55 = _dafny.MultiSet.fromElements(_111_v23, _111_v23);
      let _146_v56;
      _146_v56 = _dafny.Seq.of(_145_v55);
      (_87_globalState).f4 = (_dafny.ZERO).minus(new BigNumber(((_146_v56)[_module.__default.safeIndex((_dafny.ZERO).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]), new BigNumber((_146_v56).length))]).cardinality()));
      let _147_v58;
      _147_v58 = new _dafny.CodePoint('i'.codePointAt(0));
      let _148_v59;
      _148_v59 = _dafny.Map.Empty.slice().updateUnsafe((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))],function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Set.fromElements(_147_v58)).Elements) {
          let _149_v57 = _compr_3;
          if ((_dafny.Set.fromElements(_147_v58)).contains(_149_v57)) {
            _coll3.push([_149_v57,_111_v23]);
          }
        }
        return _coll3;
      }());
      let _hi3 = _84_v0;
      for (let _150_i3 = new BigNumber((_148_v59).length); _150_i3.isLessThan(_hi3); _150_i3 = _150_i3.plus(_dafny.ONE)) {
        let _151_v60;
        let _nw17 = new _module.C0();
        _nw17.__ctor();
        _151_v60 = _nw17;
        let _152_v61;
        _152_v61 = _dafny.MultiSet.fromElements(_100_v14);
        _152_v61 = (_152_v61).Union((_152_v61).update(_100_v14, _module.__default.abs((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])));
        let _153_v62;
        _153_v62 = _module.D0.create_DC4(_module.D0.create_DC1(_85_v1, _111_v23, _111_v23, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(437)), ((_154_v0, _155_v23) => function (_156_i4) {
  return _dafny.Map.Empty.slice().updateUnsafe(_154_v0,_155_v23);
})(_84_v0, _111_v23))).length)));
        let _157_v63;
        _157_v63 = _module.D0.create_DC4(_153_v62);
        let _158_v64;
        _158_v64 = _dafny.MultiSet.fromElements(_module.D0.create_DC4(_153_v62));
        _84_v0 = (_dafny.ZERO).minus((new BigNumber(((_158_v64).Difference(_158_v64)).cardinality())).multipliedBy(new BigNumber((_145_v55).cardinality())));
        (_87_globalState).f1 = _dafny.Seq.IsProperPrefixOf(_85_v1, _85_v1);
      }
      let _159_i5;
      _159_i5 = _dafny.ZERO;
      L1: {
        while (_111_v23) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_159_i5)) {
              break L1;
            }
            _159_i5 = (_159_i5).plus(_dafny.ONE);
            (_87_globalState).f4 = _84_v0;
            (_87_globalState).f5 = !(_module.__default.fm7(_147_v58, _87_globalState)).contains(_111_v23);
            (_87_globalState).f4 = ((_dafny.ZERO).minus(_84_v0)).minus(new BigNumber(215));
            let _index28 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
            (_86_v2)[_index28] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(698)), ((_160_v58) => function (_161_i6) {
              return _160_v58;
            })(_147_v58)), _dafny.Seq.UnicodeFromString("ydsnaxu"))).length);
          }
        }
      }
      if (_111_v23) {
        let _162_v65;
        _162_v65 = _dafny.Seq.of((_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))]);
        let _163_v66;
        _163_v66 = _module.D0.create_DC0((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
        let _index29 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length));
        (_100_v14)[_index29] = (_162_v65)[_module.__default.safeIndex((_163_v66).dtor_cf0, new BigNumber((_162_v65).length))];
        (_87_globalState).f4 = (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))];
        let _164_v67;
        _164_v67 = _dafny.Seq.of(_84_v0);
        let _165_v68;
        let _nw18 = Array((new BigNumber(26)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = false;
        _nw18[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_85_v1, _147_v58);
        _nw18[(new BigNumber(2)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(3)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(4)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(5)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(6)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(7)).toNumber()] = (_111_v23) || (_111_v23);
        _nw18[(new BigNumber(8)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(9)).toNumber()] = !(_111_v23);
        _nw18[(new BigNumber(10)).toNumber()] = !(_111_v23);
        _nw18[(new BigNumber(11)).toNumber()] = ((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]).isLessThanOrEqualTo(_84_v0);
        _nw18[(new BigNumber(12)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(13)).toNumber()] = false;
        _nw18[(new BigNumber(14)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(15)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(16)).toNumber()] = _module.__default.fm1(_164_v67, _110_v22, false, _87_globalState);
        _nw18[(new BigNumber(17)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(18)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(19)).toNumber()] = !(_111_v23) || (_111_v23);
        _nw18[(new BigNumber(20)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(21)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(22)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(23)).toNumber()] = true;
        _nw18[(new BigNumber(24)).toNumber()] = _111_v23;
        _nw18[(new BigNumber(25)).toNumber()] = false;
        _165_v68 = _nw18;
        let _index30 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_165_v68).length));
        (_165_v68)[_index30] = (_111_v23) && (_111_v23);
        let _166_v69;
        _166_v69 = _dafny.Set.fromElements(_111_v23);
        let _index31 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_165_v68).length));
        (_165_v68)[_index31] = !((_dafny.Set.fromElements(_111_v23, _111_v23, true)).IsSubsetOf(_166_v69));
        let _index32 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        (_86_v2)[_index32] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_85_v1, _85_v1), _dafny.Seq.UnicodeFromString("xvxf")), _module.__default.safeIndex(new BigNumber(-810), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_85_v1, _85_v1), _dafny.Seq.UnicodeFromString("xvxf"))).length)), _147_v58)).length);
        let _index33 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length));
        (_100_v14)[_index33] = (_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))];
      } else {
        (_87_globalState).f1 = _111_v23;
        let _167_v70;
        _167_v70 = _dafny.Seq.of((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
        let _index34 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        let _rhs24 = ((_111_v23) ? (_111_v23) : ((_145_v55).IsSubsetOf((_145_v55).update(_111_v23, _module.__default.abs(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_86_v2,_84_v0)).length))))));
        let _rhs25 = (_111_v23) && (_module.__default.fm1(_167_v70, _dafny.Set.fromElements((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]), _111_v23, _87_globalState));
        let _rhs26 = _module.__default.safeModuloInt(new BigNumber(70), (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
        let _lhs26 = _87_globalState;
        let _lhs27 = _86_v2;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        _lhs26.f1 = _rhs24;
        _111_v23 = _rhs25;
        _lhs27[_lhs28] = _rhs26;
        _module.__default.m0((!(_111_v23)) === (_module.__default.fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), ((_168_v2) => function (_169_i7) {
          return (_168_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_168_v2).length))];
        })(_86_v2)), _110_v22, _module.__default.fm1(_dafny.Seq.of(new BigNumber(758)), _110_v22, _111_v23, _87_globalState), _87_globalState)), new BigNumber((_dafny.Seq.UnicodeFromString("qjjinm")).length), _167_v70, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], _87_globalState);
        (_87_globalState).f4 = _module.__default.safeDivisionInt(_module.__default.fm0(_84_v0, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], false, _87_globalState), _84_v0);
        let _170_v71;
        let _nw19 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _170_v71 = _nw19;
        let _index35 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_170_v71).length));
        (_170_v71)[_index35] = ((_module.__default.fm1(_167_v70, _110_v22, _111_v23, _87_globalState)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-850)), ((_171_v58) => function (_172_i8) {
          return _171_v58;
        })(_147_v58))) : ((_99_v13).fm2(!(_111_v23), _84_v0, _87_globalState)));
        let _173_v72;
        let _nw20 = Array((new BigNumber(29)).toNumber()).fill(false);
        _173_v72 = _nw20;
        let _index36 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_173_v72).length));
        (_173_v72)[_index36] = false;
        let _174_v73;
        _174_v73 = _dafny.Map.Empty.slice().updateUnsafe((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))],_85_v1);
        let _175_v74;
        _175_v74 = _dafny.Seq.of(_111_v23, _111_v23, !(!(((_111_v23) ? (_111_v23) : (_111_v23)))), _module.__default.fm1(_167_v70, _110_v22, _111_v23, _87_globalState));
        let _176_v75;
        _176_v75 = _dafny.MultiSet.fromElements(_86_v2);
        let _index37 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_170_v71).length));
        let _index38 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        let _index39 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_173_v72).length));
        let _rhs27 = _85_v1;
        let _rhs28 = new BigNumber(((_module.D1.create_DC5(_174_v73)).dtor_cf11).length);
        let _rhs29 = (_84_v0).plus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
        let _rhs30 = !((_175_v74)[_module.__default.safeIndex((_dafny.ZERO).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]), new BigNumber((_175_v74).length))]);
        let _rhs31 = (_176_v75).IsProperSubsetOf(_dafny.MultiSet.fromElements(_86_v2));
        let _lhs29 = _170_v71;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_170_v71).length));
        let _lhs31 = _86_v2;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length));
        let _lhs33 = _87_globalState;
        let _lhs34 = _173_v72;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_173_v72).length));
        _lhs29[_lhs30] = _rhs27;
        _lhs31[_lhs32] = _rhs28;
        _84_v0 = _rhs29;
        _lhs33.f1 = _rhs30;
        _lhs34[_lhs35] = _rhs31;
      }
      let _177_v76;
      _177_v76 = _dafny.Seq.of(_111_v23, _111_v23, false);
      let _178_v77;
      _178_v77 = _dafny.Map.Empty.slice().updateUnsafe(_111_v23,_dafny.Seq.Concat(_dafny.Seq.of(_111_v23), _177_v76));
      let _179_v78;
      _179_v78 = _module.D0.create_DC3();
      let _180_v79;
      _180_v79 = _module.D0.create_DC4(_179_v78);
      let _181_v80;
      _181_v80 = _dafny.Seq.of(_84_v0, new BigNumber((_177_v76).length));
      let _pat_let_tv2 = _178_v77;
      let _pat_let_tv3 = _178_v77;
      let _pat_let_tv4 = _177_v76;
      let _pat_let_tv5 = _178_v77;
      let _pat_let_tv6 = _178_v77;
      let _pat_let_tv7 = _178_v77;
      let _rhs32 = function (_source2) {
        if (_source2.is_DC1) {
          let _182___mcc_h0 = (_source2).cf1;
          let _183___mcc_h1 = (_source2).cf2;
          let _184___mcc_h2 = (_source2).cf3;
          let _185___mcc_h3 = (_source2).cf4;
          let _186_cf4 = _185___mcc_h3;
          let _187_cf3 = _184___mcc_h2;
          let _188_cf2 = _183___mcc_h1;
          let _189_cf1 = _182___mcc_h0;
          return (_pat_let_tv2).Merge(_pat_let_tv3);
        } else if (_source2.is_DC2) {
          let _190___mcc_h4 = (_source2).cf5;
          let _191___mcc_h5 = (_source2).cf6;
          let _192___mcc_h6 = (_source2).cf7;
          let _193___mcc_h7 = (_source2).cf8;
          let _194___mcc_h8 = (_source2).cf9;
          let _195_cf9 = _194___mcc_h8;
          let _196_cf8 = _193___mcc_h7;
          let _197_cf7 = _192___mcc_h6;
          let _198_cf6 = _191___mcc_h5;
          let _199_cf5 = _190___mcc_h4;
          return _dafny.Map.Empty.slice().updateUnsafe(false,_pat_let_tv4);
        } else if (_source2.is_DC3) {
          return _pat_let_tv5;
        } else if (_source2.is_DC0) {
          let _200___mcc_h9 = (_source2).cf0;
          let _201_cf0 = _200___mcc_h9;
          return _pat_let_tv6;
        } else {
          let _202___mcc_h10 = (_source2).cf10;
          let _203_cf10 = _202___mcc_h10;
          return (_pat_let_tv7);
        }
      }(_180_v79);
      let _rhs33 = _101_v15;
      let _rhs34 = _84_v0;
      let _rhs35 = (_99_v13).fm2(!(_module.__default.fm1(_181_v80, _module.__default.fm8(_111_v23, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], new BigNumber(652), _87_globalState), false, _87_globalState)), _84_v0, _87_globalState);
      let _rhs36 = ((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]).multipliedBy(_84_v0);
      let _lhs36 = _87_globalState;
      _178_v77 = _rhs32;
      _101_v15 = _rhs33;
      _lhs36.f4 = _rhs34;
      _85_v1 = _rhs35;
      _84_v0 = _rhs36;
      let _204_i9;
      _204_i9 = _dafny.ZERO;
      L2: {
        while (!(_111_v23)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_204_i9)) {
              break L2;
            }
            _204_i9 = (_204_i9).plus(_dafny.ONE);
            let _205_v81;
            _205_v81 = _dafny.Map.Empty.slice().updateUnsafe(!(_111_v23),_84_v0);
            _module.__default.m0(!((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]).isEqualTo((((_205_v81).contains(_111_v23)) ? ((_205_v81).get(_111_v23)) : (new BigNumber(637)))), ((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]).minus(new BigNumber(820)), _181_v80, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), ((_206_v0) => function (_207_i10) {
              return _206_v0;
            })(_84_v0))).length), _87_globalState);
            if (_111_v23) {
              _module.__default.m0((_84_v0).isLessThan((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]), (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], ((_111_v23) ? (_181_v80) : (_181_v80)), _84_v0, _87_globalState);
              (_87_globalState).f1 = _111_v23;
              (_87_globalState).f5 = (_111_v23) === ((_dafny.Set.fromElements(_84_v0)).IsDisjointFrom(_110_v22));
              let _208_v82;
              _208_v82 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("oiyf"));
              (_87_globalState).f5 = ((_dafny.Set.fromElements(_85_v1)).Union(_208_v82)).IsSubsetOf(_module.__default.fm9(_111_v23, false, _84_v0, _87_globalState));
              let _209_v83;
              _209_v83 = _dafny.Seq.of(_85_v1);
              let _210_v84;
              _210_v84 = _dafny.Seq.of((_209_v83)[_module.__default.safeIndex(_84_v0, new BigNumber((_209_v83).length))]);
              let _rhs37 = (((_208_v82).IsDisjointFrom(function () {
                let _coll4 = new _dafny.Set();
                for (const _compr_4 of (function () {
                  let _coll5 = new _dafny.Set();
                  for (const _compr_5 of (_210_v84).Elements) {
                    let _211_v85 = _compr_5;
                    if (_dafny.Seq.contains(_210_v84, _211_v85)) {
                      _coll5.add(_211_v85);
                    }
                  }
                  return _coll5;
                }()).Elements) {
                  let _212_v86 = _compr_4;
                  if ((function () {
                    let _coll6 = new _dafny.Set();
                    for (const _compr_6 of (_210_v84).Elements) {
                      let _213_v85 = _compr_6;
                      if (_dafny.Seq.contains(_210_v84, _213_v85)) {
                        _coll6.add(_213_v85);
                      }
                    }
                    return _coll6;
                  }()).contains(_212_v86)) {
                    _coll4.add(_212_v86);
                  }
                }
                return _coll4;
              }())) ? (_145_v55) : (_145_v55));
              let _rhs38 = (_205_v81).Merge((_205_v81).Merge(_dafny.Map.Empty.slice().updateUnsafe(_111_v23,(_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])));
              _145_v55 = _rhs37;
              _205_v81 = _rhs38;
            } else {
              _module.__default.m0(!(_111_v23), ((_111_v23) ? (_84_v0) : ((_dafny.ZERO).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]))), _181_v80, new BigNumber(391), _87_globalState);
              let _214_v87;
              let _nw21 = new _module.C0();
              _nw21.__ctor();
              _214_v87 = _nw21;
              let _215_v88;
              _215_v88 = _dafny.Seq.of((_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))], (_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))], _214_v87);
              _module.__default.m0(true, _module.__default.safeModuloInt(new BigNumber((_215_v88).length), _84_v0), (_181_v80), (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], _87_globalState);
              let _216_v90;
              _216_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_85_v1).length),_181_v80);
              let _217_v91;
              _217_v91 = _dafny.Map.Empty.slice().updateUnsafe((_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))],(_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
              let _218_v92;
              _218_v92 = _dafny.Seq.of(_217_v91, _217_v91);
              let _219_v93;
              _219_v93 = _dafny.Seq.of(_dafny.Seq.update(_181_v80, _module.__default.safeIndex(new BigNumber((_module.__default.fm7(_147_v58, _87_globalState)).length), new BigNumber((_181_v80).length)), new BigNumber((function () {
                let _coll7 = new _dafny.Map();
                for (const _compr_7 of _dafny.IntegerRange(new BigNumber(958), new BigNumber(-17))) {
                  let _220_v89 = _compr_7;
                  if (((new BigNumber(958)).isLessThanOrEqualTo(_220_v89)) && ((_220_v89).isLessThan(new BigNumber(-17)))) {
                    _coll7.push([(_220_v89).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]),_111_v23]);
                  }
                }
                return _coll7;
              }()).length)), _181_v80, (((_216_v90).contains(new BigNumber((_218_v92).length))) ? ((_216_v90).get(new BigNumber((_218_v92).length))) : (_181_v80)), _181_v80);
              (_87_globalState).f1 = ((_dafny.ZERO).minus((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))])).isLessThanOrEqualTo(new BigNumber(((_219_v93)[_module.__default.safeIndex(new BigNumber(992), new BigNumber((_219_v93).length))]).length));
              _module.__default.m0(_111_v23, (_84_v0).multipliedBy(_84_v0), _181_v80, (_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))], _87_globalState);
            }
            let _221_v94;
            let _nw22 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
            _221_v94 = _nw22;
            let _222_v95;
            _222_v95 = _dafny.Map.Empty.slice().updateUnsafe((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))],_111_v23);
            let _223_v96;
            _223_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(544),_222_v95);
            let _index40 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_221_v94).length));
            (_221_v94)[_index40] = _223_v96;
            let _224_v98;
            _224_v98 = _module.D4.create_DC10(function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-558), new BigNumber(248))) {
    let _225_v97 = _compr_8;
    if (((new BigNumber(-558)).isLessThanOrEqualTo(_225_v97)) && ((_225_v97).isLessThan(new BigNumber(248)))) {
      _coll8.push([(_225_v97).multipliedBy(_84_v0),_dafny.Map.Empty.slice().updateUnsafe(_84_v0,_111_v23)]);
    }
  }
  return _coll8;
}());
            let _index41 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_221_v94).length));
            (_221_v94)[_index41] = ((_224_v98).dtor_cf15).Merge(_223_v96);
            let _226_v99;
            _226_v99 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_111_v23,(_100_v14)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_100_v14).length))]));
            let _227_v100;
            _227_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_85_v1).length),_226_v99);
            let _228_v101;
            _228_v101 = _dafny.Seq.of(_226_v99, _226_v99, (((_227_v100).contains(new BigNumber(921))) ? ((_227_v100).get(new BigNumber(921))) : (_226_v99)));
            let _229_v102;
            let _nw23 = Array((new BigNumber(24)).toNumber());
            _nw23[(_dafny.ZERO).toNumber()] = (_145_v55).contains(_111_v23);
            _nw23[(_dafny.ONE).toNumber()] = true;
            _nw23[(new BigNumber(2)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(3)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(4)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(5)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(6)).toNumber()] = _module.__default.fm1(_181_v80, _110_v22, _111_v23, _87_globalState);
            _nw23[(new BigNumber(7)).toNumber()] = (new BigNumber(407)).isLessThan((_86_v2)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_86_v2).length))]);
            _nw23[(new BigNumber(8)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(9)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(10)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(11)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(12)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(13)).toNumber()] = !_dafny.Seq.contains(_181_v80, _84_v0);
            _nw23[(new BigNumber(14)).toNumber()] = !(_111_v23) || (_111_v23);
            _nw23[(new BigNumber(15)).toNumber()] = (_226_v99).IsProperSubsetOf((_228_v101)[_module.__default.safeIndex(_84_v0, new BigNumber((_228_v101).length))]);
            _nw23[(new BigNumber(16)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(17)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(18)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(19)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(20)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(21)).toNumber()] = !(_111_v23);
            _nw23[(new BigNumber(22)).toNumber()] = _111_v23;
            _nw23[(new BigNumber(23)).toNumber()] = _111_v23;
            _229_v102 = _nw23;
            _229_v102 = _229_v102;
          }
        }
      }
      let _230_v103;
      let _nw24 = Array((new BigNumber(20)).toNumber()).fill(false);
      _230_v103 = _nw24;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_230_v103).length))) {
        let _231_i11 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_231_i11)) && ((_231_i11).isLessThan(new BigNumber((_230_v103).length))))) {
          (_230_v103)[(_231_i11)] = _111_v23;
        }
      }
      let _232_v104;
      _232_v104 = _module.D4.create_DC12(_147_v58, _109_v21, _111_v23, new BigNumber((_85_v1).length), _99_v13);
      let _233_v105;
      _233_v105 = _dafny.MultiSet.fromElements((_232_v104).dtor_cf20);
      _233_v105 = _233_v105;
      process.stdout.write(_dafny.toString(_84_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_85_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_87_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_87_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_87_globalState).f2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_87_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_87_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_87_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_101_v15).dtor_cf1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v15).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v15).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v15).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_107_v20).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0))),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_109_v21).equals(_dafny.MultiSet.fromElements(new BigNumber(624)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_110_v22).equals(_dafny.Set.fromElements(new BigNumber(-744), new BigNumber(2), _dafny.ONE, new BigNumber(-1), new BigNumber(624)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_111_v23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_v55).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_146_v56, _dafny.Seq.of(_dafny.MultiSet.fromElements(false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_v58));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_v59).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(624),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_177_v76, _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v77).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false, false, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_181_v80, _dafny.Seq.of(new BigNumber(70), new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_204_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v103)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_232_v104).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_232_v104).dtor_cf17).equals(_dafny.MultiSet.fromElements(new BigNumber(624)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_232_v104).dtor_cf18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_232_v104).dtor_cf19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_233_v105).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC2(cf5, cf6, cf7, cf8, cf9) {
      let $dt = new D0(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC3() {
      let $dt = new D0(2);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC4(cf10) {
      let $dt = new D0(4);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get is_DC4() { return this.$tag === 4; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + this.cf1.toVerbatimString(true) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + this.cf5.toVerbatimString(true) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 4) {
        return "D0.DC4" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.Seq.UnicodeFromString(""), false, false, _dafny.ZERO);
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
    static create_DC6(cf12) {
      let $dt = new D1(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7() {
      let $dt = new D1(1);
      return $dt;
    }
    static create_DC5(cf11) {
      let $dt = new D1(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC7";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC6(_dafny.ZERO);
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
    static create_DC8(cf13) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
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
    static create_DC9(cf14) {
      let $dt = new D3(0);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14);
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC12(cf16, cf17, cf18, cf19, cf20) {
      let $dt = new D4(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC10(cf15) {
      let $dt = new D4(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf15) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11();
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
    static create_DC13(cf21) {
      let $dt = new D5(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf21) + ")";
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

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = false;
      this.f4 = _dafny.ZERO;
      this.f5 = false;
      this._f0 = false;
      this._f2 = [];
      this._f3 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.UnicodeFromString("ytnl")) : (_dafny.Seq.UnicodeFromString("wfgpisr"))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qsbhxkh"), _dafny.Seq.UnicodeFromString("p")));
    };
    fm3(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),(new BigNumber(376)).multipliedBy(new BigNumber(-205)));
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
