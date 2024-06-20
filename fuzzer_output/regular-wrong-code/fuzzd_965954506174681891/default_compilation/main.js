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
      return ((new BigNumber((_dafny.Seq.of(true, !(false))).length)).multipliedBy(new BigNumber(840))).isLessThan((new BigNumber(-513)).plus(new BigNumber(882)));
    };
    static fm1(p0, p1, p2, globalState) {
      return new BigNumber(736);
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("euwd"), _dafny.Seq.UnicodeFromString("d"));
    };
    static fm5(p0, globalState) {
      return function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), function (_0_i0) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-211)), function (_1_i1) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("bmpbyeb"), _dafny.Seq.UnicodeFromString("jeeg"))).Elements) {
          let _2_v0 = _compr_0;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), function (_0_i0) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-211)), function (_1_i1) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("bmpbyeb"), _dafny.Seq.UnicodeFromString("jeeg"))).contains(_2_v0)) {
            _coll0.add(_2_v0);
          }
        }
        return _coll0;
      }();
    };
    static fm6(p0, p1, p2, globalState) {
      return _module.D1.create_DC4((function () {
  let _coll1 = new _dafny.Set();
  for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(240)), function (_3_i0) {
    return _dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), function (_4_i1) {
      return new _dafny.CodePoint('m'.codePointAt(0));
    });
  })).Elements) {
    let _5_v0 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(240)), function (_3_i0) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), function (_4_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      });
    }), _5_v0)) {
      _coll1.add(_5_v0);
    }
  }
  return _coll1;
}()).Intersect(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("khuk"), _dafny.Seq.UnicodeFromString("gb"), _dafny.Seq.UnicodeFromString("gbthcxlck"))));
    };
    static fm7(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(582), (_module.D7.create_DC17(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(121),new BigNumber(901))).length), false, true)).dtor_cf29, new BigNumber(-450)), _dafny.Seq.of(new BigNumber(630), new BigNumber(714))), _dafny.Seq.of(new BigNumber(549)));
    };
    static fm8(p0, p1, globalState) {
      return _dafny.Seq.of(!(new BigNumber(528)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),(_module.D7.create_DC17(true, new BigNumber((_dafny.Seq.of(new BigNumber(706), new BigNumber(932))).length), true, !(true))).dtor_cf30)).length)), (new BigNumber(561)).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(469), new BigNumber(-47))) {
          let _6_v0 = _compr_2;
          if (((new BigNumber(469)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(-47)))) {
            _coll2.add((_6_v0).minus(new BigNumber(437)));
          }
        }
        return _coll2;
      }()).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).length), new BigNumber(581), (_dafny.ZERO).minus(new BigNumber(-48)))).cardinality())), false, !((new BigNumber((_dafny.Seq.of(_dafny.Seq.of(!(true), !(true)), _dafny.Seq.of(true, true, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(!(true), true, true), _dafny.Seq.of(true, false, true))).length)).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()))), false);
    };
    static fm9(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm10(p0, p1, globalState) {
      if ((true) || (!(true))) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(65)), function (_7_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("xm"));
    };
    static fm12(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(true, false))).Union((_dafny.Set.fromElements(true, false)).Difference(_dafny.Set.fromElements(true)));
    };
    static fm13(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(986),new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length))).length), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).length)))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-994), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-991))).length))).length))),_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-577),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(660),new BigNumber(407)))).length), new BigNumber(584)));
    };
    static m0(globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _8_v0;
      _8_v0 = new BigNumber(301);
      let _9_v1;
      _9_v1 = false;
      let _10_v2;
      _10_v2 = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,new _dafny.CodePoint('l'.codePointAt(0)));
      let _hi0 = (_dafny.ZERO).minus(new BigNumber(((_10_v2).Merge(_dafny.Map.Empty.slice().updateUnsafe(_9_v1,new _dafny.CodePoint('k'.codePointAt(0))))).length));
      for (let _11_i0 = _8_v0; _11_i0.isLessThan(_hi0); _11_i0 = _11_i0.plus(_dafny.ONE)) {
        r1 = _8_v0;
        _8_v0 = _8_v0;
        if (_9_v1) {
          let _12_v3;
          _12_v3 = _dafny.Seq.UnicodeFromString("r");
          let _13_v4;
          _13_v4 = new _dafny.CodePoint('l'.codePointAt(0));
          let _rhs0 = _dafny.Seq.update(_dafny.Seq.of(_9_v1, _dafny.Seq.IsProperPrefixOf(_12_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(196)), function (_14_i1) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })), _9_v1, _9_v1, (_9_v1) && (_module.__default.fm0(_9_v1, _9_v1, _12_v3, globalState))), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(267)), ((_15_i0) => function (_16_i2) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-835)), ((_17_i0) => function (_18_i3) {
              return _17_i0;
            })(_15_i0));
          })(_11_i0))).length), new BigNumber((_dafny.Seq.of(_9_v1, _dafny.Seq.IsProperPrefixOf(_12_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(196)), function (_19_i1) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })), _9_v1, _9_v1, (_9_v1) && (_module.__default.fm0(_9_v1, _9_v1, _12_v3, globalState)))).length)), (_9_v1) && (_9_v1));
          let _rhs1 = _11_i0;
          let _rhs2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("quhelj"), _dafny.Seq.update(_12_v3, _module.__default.safeIndex(_8_v0, new BigNumber((_12_v3).length)), _13_v4)), _12_v3);
          let _lhs0 = globalState;
          let _lhs1 = globalState;
          _lhs0.f9 = _rhs0;
          _lhs1.f7 = _rhs1;
          _12_v3 = _rhs2;
          let _20_v5;
          let _init0 = ((_21_i0) => function (_22_i4) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_21_i0), _dafny.Seq.of(_21_i0));
          })(_11_i0);
          let _nw0 = Array((new BigNumber(27)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
            _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _20_v5 = _nw0;
          let _23_v6;
          let _nw1 = Array((new BigNumber(26)).toNumber()).fill(false);
          _23_v6 = _nw1;
          let _24_v7;
          let _nw2 = Array((new BigNumber(9)).toNumber()).fill([]);
          _24_v7 = _nw2;
          let _index0 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_24_v7).length));
          (_24_v7)[_index0] = _23_v6;
          let _25_v8;
          _25_v8 = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,_23_v6);
          let _26_v9;
          _26_v9 = _module.D5.create_DC11(_9_v1, _11_i0, true, _23_v6, new BigNumber(319));
          let _27_v10;
          _27_v10 = _dafny.Map.Empty.slice().updateUnsafe((_12_v3)[_module.__default.safeIndex(_8_v0, new BigNumber((_12_v3).length))],_12_v3);
          let _28_v11;
          _28_v11 = _dafny.Set.fromElements(_module.__default.fm1(false, _8_v0, (_8_v0), globalState));
          let _index1 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_24_v7).length));
          let _rhs3 = _20_v5;
          let _rhs4 = (((_25_v8).contains(!(_9_v1))) ? ((_25_v8).get(!(_9_v1))) : (_23_v6));
          let _rhs5 = (_26_v9).dtor_cf14;
          let _rhs6 = (_module.__default.safeModuloInt(_8_v0, new BigNumber((_28_v11).length))).isLessThan(new BigNumber((_27_v10).length));
          let _lhs2 = _24_v7;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_24_v7).length));
          _20_v5 = _rhs3;
          _23_v6 = _rhs4;
          _lhs2[_lhs3] = _rhs5;
          _9_v1 = _rhs6;
          _8_v0 = _11_i0;
          let _29_v12;
          let _nw3 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _29_v12 = _nw3;
          let _30_v13;
          _30_v13 = _dafny.Seq.of(_8_v0);
          let _31_v14;
          _31_v14 = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,_9_v1);
          let _32_v15;
          let _nw4 = new _module.C0();
          _nw4.__ctor(_9_v1, _31_v14);
          _32_v15 = _nw4;
          let _33_v16;
          _33_v16 = _dafny.Seq.of(new BigNumber((_30_v13).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_32_v15,_8_v0)).length));
          let _34_v17;
          _34_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_30_v13).length),_9_v1);
          let _35_v18;
          _35_v18 = _dafny.MultiSet.fromElements(_11_i0, _8_v0);
          let _36_v19;
          _36_v19 = _dafny.Seq.of(_9_v1, (((_34_v17).contains((((_35_v18).contains((_dafny.ZERO).minus(_8_v0))) ? ((_35_v18).get((_dafny.ZERO).minus(_8_v0))) : (_8_v0)))) ? ((_34_v17).get((((_35_v18).contains((_dafny.ZERO).minus(_8_v0))) ? ((_35_v18).get((_dafny.ZERO).minus(_8_v0))) : (_8_v0)))) : (_32_v15.f19)), _32_v15.f19);
          let _37_v20;
          _37_v20 = _dafny.Map.Empty.slice().updateUnsafe(_29_v12,_36_v19);
          _37_v20 = _37_v20;
          _13_v4 = _13_v4;
        } else {
          (globalState).f8 = _8_v0;
          let _rhs7 = (_8_v0).plus(_8_v0);
          let _rhs8 = _9_v1;
          r1 = _rhs7;
          _9_v1 = _rhs8;
          let _38_v21;
          let _nw5 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _38_v21 = _nw5;
          _38_v21 = _38_v21;
          let _39_v22;
          _39_v22 = _dafny.Seq.UnicodeFromString("xbhg");
          let _40_v23;
          let _nw6 = new _module.C0();
          _nw6.__ctor(_9_v1, _module.__default.fm9(_8_v0, new BigNumber((_39_v22).length), _module.__default.fm1(_9_v1, _11_i0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("q"),_8_v0)).length), globalState), globalState));
          _40_v23 = _nw6;
          let _41_v24;
          _41_v24 = new _dafny.CodePoint('y'.codePointAt(0));
          let _42_v25;
          _42_v25 = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,new BigNumber(770));
          let _rhs9 = _module.__default.fm0(_40_v23.f19, (_8_v0).isLessThanOrEqualTo(_8_v0), _dafny.Seq.update(_39_v22, _module.__default.safeIndex(_11_i0, new BigNumber((_39_v22).length)), _41_v24), globalState);
          let _rhs10 = _40_v23;
          let _rhs11 = (_dafny.ZERO).minus((_11_i0).plus(new BigNumber((_42_v25).length)));
          let _rhs12 = _39_v22;
          let _rhs13 = ((_9_v1) ? (_40_v23.f19) : (_9_v1));
          _9_v1 = _rhs9;
          _40_v23 = _rhs10;
          r0 = _rhs11;
          _39_v22 = _rhs12;
          _9_v1 = _rhs13;
          _8_v0 = _11_i0;
        }
        let _43_v26;
        let _nw7 = Array((new BigNumber(2)).toNumber()).fill(false);
        _43_v26 = _nw7;
        let _44_v27;
        _44_v27 = _dafny.Seq.of(_43_v26, _43_v26, ((_9_v1) ? (_43_v26) : (_43_v26)));
        _43_v26 = (_44_v27)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_44_v27).length))];
      }
      r0 = _8_v0;
      let _45_v28;
      _45_v28 = new _dafny.CodePoint('f'.codePointAt(0));
      let _46_v29;
      _46_v29 = _dafny.Seq.of(_45_v28, _45_v28, _45_v28);
      let _47_v30;
      _47_v30 = _dafny.Seq.of(_dafny.Seq.IsProperPrefixOf(_46_v29, _dafny.Seq.Create(_module.__default.abs(new BigNumber(373)), ((_48_v28) => function (_49_i5) {
        return _48_v28;
      })(_45_v28))), _9_v1);
      let _50_v31;
      _50_v31 = _dafny.Map.Empty.slice().updateUnsafe(_8_v0,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), ((_51_v0) => function (_52_i6) {
        return _51_v0;
      })(_8_v0))).length));
      let _rhs14 = (((!(_9_v1)) ? (new BigNumber(12)) : (_8_v0))).multipliedBy(_8_v0);
      let _rhs15 = (_47_v30)[_module.__default.safeIndex((_8_v0).multipliedBy((((_50_v31).contains(new BigNumber((_46_v29).length))) ? ((_50_v31).get(new BigNumber((_46_v29).length))) : (_module.__default.fm1(_9_v1, _8_v0, _8_v0, globalState)))), new BigNumber((_47_v30).length))];
      r1 = _rhs14;
      _9_v1 = _rhs15;
      let _53_v32;
      let _nw8 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _53_v32 = _nw8;
      let _54_v33;
      _54_v33 = _46_v29;
      let _index2 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_53_v32).length));
      (_53_v32)[_index2] = (_54_v33);
      let _index3 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_53_v32).length));
      (_53_v32)[_index3] = _46_v29;
      let _55_v34;
      let _init1 = ((_56_v0) => function (_57_i7) {
        return (_56_v0).isLessThan(_56_v0);
      })(_8_v0);
      let _nw9 = Array((new BigNumber(27)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw9.length); _i0_1++) {
        _nw9[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _55_v34 = _nw9;
      let _index4 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
      (_55_v34)[_index4] = (_8_v0).isLessThan(new BigNumber(483));
      let _58_v35;
      let _nw10 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _58_v35 = _nw10;
      let _59_v36;
      _59_v36 = _dafny.Map.Empty.slice().updateUnsafe(_47_v30,_58_v35);
      let _60_v37;
      _60_v37 = _module.D0.create_DC0((_59_v36).Merge(_59_v36));
      let _pat_let_tv0 = _59_v36;
      let _index5 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
      let _rhs16 = _9_v1;
      let _rhs17 = _8_v0;
      let _rhs18 = _9_v1;
      let _rhs19 = function (_pat_let0_0) {
        return function (_61_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_62_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_62_dt__update_hcf0_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_60_v37);
      let _lhs4 = _55_v34;
      let _lhs5 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
      _9_v1 = _rhs16;
      r1 = _rhs17;
      _lhs4[_lhs5] = _rhs18;
      _60_v37 = _rhs19;
      if ((_8_v0).isLessThan(_8_v0)) {
        let _index6 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
        (_55_v34)[_index6] = !(_9_v1);
        r0 = _8_v0;
        (globalState).f7 = _8_v0;
        if (_9_v1) {
          let _63_v38;
          _63_v38 = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,!(_9_v1));
          let _64_v39;
          _64_v39 = _dafny.Seq.of(_63_v38, _63_v38, _63_v38, _63_v38, (_63_v38).update((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))], _9_v1));
          let _65_v40;
          _65_v40 = _dafny.Seq.of(_module.__default.fm1(true, new BigNumber((_50_v31).length), _8_v0, globalState), _8_v0);
          let _66_v41;
          let _nw11 = Array((new BigNumber(21)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = _63_v38;
          _nw11[(_dafny.ONE).toNumber()] = _63_v38;
          _nw11[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,(_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]);
          _nw11[(new BigNumber(3)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(4)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(5)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(6)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(7)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(8)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))],(_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]);
          _nw11[(new BigNumber(10)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,(_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]);
          _nw11[(new BigNumber(12)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(13)).toNumber()] = (_64_v39)[_module.__default.safeIndex(new BigNumber((_65_v40).length), new BigNumber((_64_v39).length))];
          _nw11[(new BigNumber(14)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,(_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]);
          _nw11[(new BigNumber(16)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(17)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))],_9_v1);
          _nw11[(new BigNumber(19)).toNumber()] = _63_v38;
          _nw11[(new BigNumber(20)).toNumber()] = _63_v38;
          _66_v41 = _nw11;
          let _67_v42;
          _67_v42 = _dafny.Map.Empty.slice().updateUnsafe(false,_66_v41);
          let _68_v43;
          _68_v43 = _dafny.Seq.of(_66_v41, _66_v41, _66_v41);
          let _69_v44;
          _69_v44 = _dafny.Set.fromElements(_8_v0, new BigNumber((_65_v40).length));
          let _70_v45;
          _70_v45 = _9_v1;
          let _71_v46;
          let _nw12 = Array((new BigNumber(27)).toNumber());
          _nw12[(_dafny.ZERO).toNumber()] = _66_v41;
          _nw12[(_dafny.ONE).toNumber()] = _66_v41;
          _nw12[(new BigNumber(2)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(3)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(4)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(5)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(6)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(7)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(8)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(9)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(10)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(11)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(12)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(13)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(14)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(15)).toNumber()] = (((_67_v42).contains((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))])) ? ((_67_v42).get((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))])) : (_66_v41));
          _nw12[(new BigNumber(16)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(17)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(18)).toNumber()] = (_68_v43)[_module.__default.safeIndex(new BigNumber((_69_v44).length), new BigNumber((_68_v43).length))];
          _nw12[(new BigNumber(19)).toNumber()] = ((_9_v1) ? (_66_v41) : (_66_v41));
          _nw12[(new BigNumber(20)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(21)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(22)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(23)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(24)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(25)).toNumber()] = _66_v41;
          _nw12[(new BigNumber(26)).toNumber()] = (((_70_v45)) ? (_66_v41) : (_66_v41));
          _71_v46 = _nw12;
          let _index7 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_71_v46).length));
          (_71_v46)[_index7] = _66_v41;
          let _index8 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_71_v46).length));
          let _init2 = ((_72_v38) => function (_73_i8) {
            return _72_v38;
          })(_63_v38);
          let _nw13 = Array((new BigNumber(23)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw13.length); _i0_2++) {
            _nw13[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          (_71_v46)[_index8] = _nw13;
          let _74_v47;
          let _nw14 = new _module.C0();
          _nw14.__ctor(_9_v1, _63_v38);
          _74_v47 = _nw14;
          _74_v47 = _74_v47;
          r1 = _8_v0;
          _9_v1 = !((_8_v0).minus(new BigNumber((_63_v38).length))).isEqualTo(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_46_v29, _module.__default.safeIndex(_8_v0, new BigNumber((_46_v29).length)), _45_v28)).length))).minus((_dafny.ZERO).minus(_8_v0)));
          (globalState).f8 = _8_v0;
        } else {
          _8_v0 = (_dafny.ZERO).minus(_8_v0);
          let _75_v48;
          _75_v48 = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,_9_v1);
          let _76_v49;
          let _nw15 = new _module.C0();
          _nw15.__ctor(_9_v1, _75_v48);
          _76_v49 = _nw15;
          let _77_v50;
          _77_v50 = _dafny.Seq.of(_58_v35);
          let _78_v51;
          _78_v51 = _dafny.MultiSet.fromElements(_58_v35, (_77_v50)[_module.__default.safeIndex(new BigNumber((_46_v29).length), new BigNumber((_77_v50).length))], _58_v35, _58_v35);
          let _79_v52;
          _79_v52 = _dafny.Map.Empty.slice().updateUnsafe((_8_v0).minus(_8_v0),_76_v49);
          let _80_v53;
          let _nw16 = Array((new BigNumber(5)).toNumber()).fill(_module.D1.Default());
          _80_v53 = _nw16;
          let _81_v54;
          _81_v54 = _module.D1.create_DC4(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("qvldbtq"), _dafny.Seq.UnicodeFromString("aiws")));
          let _82_v55;
          _82_v55 = _module.D1.create_DC5(_81_v54);
          let _index9 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_80_v53).length));
          (_80_v53)[_index9] = _82_v55;
          let _83_v56;
          _83_v56 = _dafny.MultiSet.fromElements(_47_v30);
          let _index10 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_80_v53).length));
          let _rhs20 = ((!(_9_v1)) ? (_76_v49) : (_76_v49));
          let _rhs21 = _78_v51;
          let _rhs22 = ((_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(_47_v30))).Intersect(_83_v56)).cardinality()))).plus(new BigNumber(79));
          let _rhs23 = (_79_v52).Merge(_79_v52);
          let _rhs24 = _82_v55;
          let _lhs6 = globalState;
          let _lhs7 = _80_v53;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_80_v53).length));
          _76_v49 = _rhs20;
          _78_v51 = _rhs21;
          _lhs6.f7 = _rhs22;
          _79_v52 = _rhs23;
          _lhs7[_lhs8] = _rhs24;
          _10_v2 = (_10_v2).update(_9_v1, _45_v28);
          let _index11 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
          (_55_v34)[_index11] = _dafny.Seq.contains(_47_v30, _9_v1);
          let _84_v57;
          _84_v57 = _dafny.Map.Empty.slice().updateUnsafe(_45_v28,_76_v49.f19);
          _84_v57 = (_84_v57).Merge(_84_v57);
        }
        r0 = (_dafny.ZERO).minus(_8_v0);
      } else {
        let _index12 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length));
        (_58_v35)[_index12] = _8_v0;
        let _index13 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length));
        (_58_v35)[_index13] = _8_v0;
        let _85_v58;
        _85_v58 = _dafny.Set.fromElements(_dafny.Seq.of((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]), _module.__default.fm8(new BigNumber(-7), _9_v1, globalState));
        let _index14 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
        (_55_v34)[_index14] = (_dafny.Set.fromElements(_47_v30)).IsProperSubsetOf(_85_v58);
        _8_v0 = _8_v0;
        let _86_v59;
        _86_v59 = _dafny.Set.fromElements(_8_v0, (_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]);
        let _87_v60;
        _87_v60 = _dafny.MultiSet.fromElements((_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))], (_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]);
        let _88_v61;
        _88_v61 = _dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC14(_87_v60)).dtor_cf18,(_53_v32)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_53_v32).length))]);
        let _89_v62;
        _89_v62 = _module.D5.create_DC10(_86_v59, _88_v61);
        let _source0 = _89_v62;
        if (_source0.is_DC10) {
          let _90___mcc_h0 = (_source0).cf9;
          let _91___mcc_h1 = (_source0).cf10;
          let _92_cf10 = _91___mcc_h1;
          let _93_cf9 = _90___mcc_h0;
          let _index15 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length));
          (_58_v35)[_index15] = _8_v0;
          let _94_v63;
          _94_v63 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(318)), ((_95_v28) => function (_96_i9) {
            return _95_v28;
          })(_45_v28)));
          _9_v1 = (_94_v63).IsProperSubsetOf((_94_v63).Union(_94_v63));
          let _97_v64;
          let _nw17 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
          _97_v64 = _nw17;
          let _98_v65;
          _98_v65 = _module.D1.create_DC3(_97_v64);
          let _pat_let_tv1 = _97_v64;
          let _99_v66;
          let _nw18 = Array((new BigNumber(28)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _98_v65;
          _nw18[(_dafny.ONE).toNumber()] = _98_v65;
          _nw18[(new BigNumber(2)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(3)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(4)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(5)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(6)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(7)).toNumber()] = _module.D1.create_DC3(_97_v64);
          _nw18[(new BigNumber(8)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(9)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(10)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(11)).toNumber()] = ((_9_v1) ? (_98_v65) : (_98_v65));
          _nw18[(new BigNumber(12)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(13)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(14)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(15)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(16)).toNumber()] = _module.D1.create_DC3((_module.D1.create_DC3(_97_v64)).dtor_cf2);
          _nw18[(new BigNumber(17)).toNumber()] = ((_9_v1) ? (_98_v65) : (_98_v65));
          _nw18[(new BigNumber(18)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(19)).toNumber()] = _module.D1.create_DC3(_97_v64);
          _nw18[(new BigNumber(20)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(21)).toNumber()] = _module.D1.create_DC3(_97_v64);
          _nw18[(new BigNumber(22)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(23)).toNumber()] = function (_pat_let2_0) {
            return function (_100_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_101_dt__update_hcf2_h0) {
                  return _module.D1.create_DC3(_101_dt__update_hcf2_h0);
                }(_pat_let3_0);
              }(_pat_let_tv1);
            }(_pat_let2_0);
          }(_98_v65);
          _nw18[(new BigNumber(24)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(25)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(26)).toNumber()] = _98_v65;
          _nw18[(new BigNumber(27)).toNumber()] = _98_v65;
          _99_v66 = _nw18;
          let _index16 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_99_v66).length));
          (_99_v66)[_index16] = _98_v65;
          let _index17 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_99_v66).length));
          (_99_v66)[_index17] = _98_v65;
          let _102_v67;
          _102_v67 = _module.D7.create_DC14(_87_v60);
          let _index18 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length));
          let _rhs25 = ((_module.__default.fm0(!(!(_9_v1)), _9_v1, _dafny.Seq.UnicodeFromString("i"), globalState)) ? (_102_v67) : (_102_v67));
          let _rhs26 = (_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))];
          let _lhs9 = _58_v35;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length));
          _102_v67 = _rhs25;
          _lhs9[_lhs10] = _rhs26;
        } else if (_source0.is_DC11) {
          let _103___mcc_h2 = (_source0).cf11;
          let _104___mcc_h3 = (_source0).cf12;
          let _105___mcc_h4 = (_source0).cf13;
          let _106___mcc_h5 = (_source0).cf14;
          let _107___mcc_h6 = (_source0).cf15;
          let _108_cf15 = _107___mcc_h6;
          let _109_cf14 = _106___mcc_h5;
          let _110_cf13 = _105___mcc_h4;
          let _111_cf12 = _104___mcc_h3;
          let _112_cf11 = _103___mcc_h2;
          let _113_v68;
          _113_v68 = _module.D5.create_DC11((_47_v30)[_module.__default.safeIndex(_111_cf12, new BigNumber((_47_v30).length))], _8_v0, _9_v1, _109_cf14, (_dafny.ZERO).minus(_111_cf12));
          _113_v68 = _113_v68;
          _8_v0 = ((_112_cf11) ? (_8_v0) : (_8_v0));
          _112_cf11 = _112_cf11;
          let _114_v69;
          let _nw19 = new _module.C0();
          _nw19.__ctor(((_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]).isLessThan(_8_v0), _dafny.Map.Empty.slice().updateUnsafe((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))],_9_v1));
          _114_v69 = _nw19;
          _114_v69 = _114_v69;
        } else if (_source0.is_DC9) {
          let _115___mcc_h7 = (_source0).cf8;
          let _116_cf8 = _115___mcc_h7;
          let _117_v70;
          _117_v70 = new BigNumber(96);
          let _118_v71;
          _118_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(_9_v1),(_117_v70));
          let _119_v72;
          let _nw20 = Array((new BigNumber(20)).toNumber()).fill(_module.D5.Default());
          _119_v72 = _nw20;
          let _120_v73;
          _120_v73 = _dafny.MultiSet.fromElements(_119_v72);
          _118_v71 = (_118_v71).update((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))], _module.__default.safeDivisionInt(new BigNumber((_120_v73).cardinality()), new BigNumber(((_53_v32)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_53_v32).length))]).length)));
          let _121_v74;
          _121_v74 = _module.D5.create_DC11(_116_cf8.f19, (_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))], (_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))], _55_v34, (_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]);
          let _122_v75;
          _122_v75 = _dafny.MultiSet.fromElements((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]);
          let _123_v76;
          _123_v76 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(_116_cf8.f19, _9_v1)), _122_v75, _122_v75);
          let _124_v77;
          let _nw21 = Array((new BigNumber(11)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _8_v0;
          _nw21[(_dafny.ONE).toNumber()] = ((_121_v74).dtor_cf15).plus(new BigNumber(((_116_cf8).f20).length));
          _nw21[(new BigNumber(2)).toNumber()] = _8_v0;
          _nw21[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ivkit")).length);
          _nw21[(new BigNumber(4)).toNumber()] = (_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))];
          _nw21[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]);
          _nw21[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_8_v0).plus((_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]));
          _nw21[(new BigNumber(7)).toNumber()] = _8_v0;
          _nw21[(new BigNumber(8)).toNumber()] = new BigNumber((_123_v76).length);
          _nw21[(new BigNumber(9)).toNumber()] = _8_v0;
          _nw21[(new BigNumber(10)).toNumber()] = _8_v0;
          _124_v77 = _nw21;
          _124_v77 = _124_v77;
          let _125_v78;
          _125_v78 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("pbduko"));
          _125_v78 = (_125_v78).Union(_125_v78);
          let _126_v79;
          let _init3 = ((_127_v30) => function (_128_i10) {
            return _127_v30;
          })(_47_v30);
          let _nw22 = Array((new BigNumber(26)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw22.length); _i0_3++) {
            _nw22[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _126_v79 = _nw22;
          let _129_v80;
          _129_v80 = _dafny.Map.Empty.slice().updateUnsafe(_126_v79,(_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]);
          _129_v80 = (_129_v80).update(_126_v79, new BigNumber((_dafny.MultiSet.fromElements((_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))])).cardinality()));
        } else {
          let _130___mcc_h8 = (_source0).cf16;
          let _131_cf16 = _130___mcc_h8;
          let _132_v81;
          _132_v81 = _dafny.Map.Empty.slice().updateUnsafe(_9_v1,!((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]));
          let _index19 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
          let _rhs27 = (_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))];
          let _rhs28 = new BigNumber((_132_v81).length);
          let _rhs29 = true;
          let _lhs11 = globalState;
          let _lhs12 = _55_v34;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length));
          _9_v1 = _rhs27;
          _lhs11.f8 = _rhs28;
          _lhs12[_lhs13] = _rhs29;
          let _133_v82;
          let _nw23 = Array((new BigNumber(4)).toNumber());
          _133_v82 = _nw23;
          _133_v82 = _133_v82;
          (globalState).f8 = _module.__default.safeModuloInt(new BigNumber(-109), (_dafny.ZERO).minus((_58_v35)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_58_v35).length))]));
          let _134_v83;
          let _nw24 = new _module.C0();
          _nw24.__ctor((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))], (_132_v81).update(false, (_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))]));
          _134_v83 = _nw24;
        }
        let _135_v84;
        _135_v84 = _dafny.Set.fromElements(_9_v1);
        let _136_v85;
        _136_v85 = _dafny.Map.Empty.slice().updateUnsafe((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))],_135_v84);
        let _137_v86;
        _137_v86 = _dafny.Seq.of((((_136_v85).contains((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))])) ? ((_136_v85).get((_55_v34)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_55_v34).length))])) : (_135_v84)), _135_v84);
        _137_v86 = _137_v86;
      }
      r0 = _8_v0;
      r1 = _module.__default.safeModuloInt(_8_v0, new BigNumber(554));
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _138_v0;
      _138_v0 = new _dafny.CodePoint('u'.codePointAt(0));
      let _139_v1;
      let _init4 = function (_140_i0) {
        return _module.__default.safeDivisionInt(_140_i0, new BigNumber((_dafny.Seq.UnicodeFromString("y")).length));
      };
      let _nw25 = Array((new BigNumber(13)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw25.length); _i0_4++) {
        _nw25[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _139_v1 = _nw25;
      let _141_v2;
      _141_v2 = _dafny.Map.Empty.slice().updateUnsafe(_138_v0,_139_v1);
      let _142_v3;
      _142_v3 = false;
      let _143_v4;
      _143_v4 = _dafny.Set.fromElements(_142_v3, _142_v3);
      let _144_v5;
      _144_v5 = new BigNumber(172);
      let _145_v6;
      _145_v6 = _dafny.Seq.of(_142_v3, _142_v3, _142_v3);
      let _146_v7;
      _146_v7 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,false);
      let _147_globalState;
      let _nw26 = new _module.GlobalState();
      _nw26.__ctor(true, _dafny.Seq.UnicodeFromString("pvfmfbioj"), _141_v2, _143_v4, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(492)), ((_148_v0) => function (_149_i1) {
        return _148_v0;
      })(_138_v0)),_144_v5), new BigNumber(-984), false, new BigNumber(148), new BigNumber(-376), _dafny.Seq.Concat(_145_v6, _145_v6), false, _139_v1, true, _146_v7, new BigNumber(363), new BigNumber(680), true, false, false);
      _147_globalState = _nw26;
      let _150_v8;
      _150_v8 = _dafny.MultiSet.fromElements(_142_v3);
      _142_v3 = !((_dafny.MultiSet.fromElements(_142_v3, _142_v3, _142_v3)).Intersect(_150_v8)).contains(_142_v3);
      let _151_v9;
      _151_v9 = _dafny.Seq.UnicodeFromString("qu");
      let _152_i2;
      _152_i2 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm0(_142_v3, _142_v3, _151_v9, _147_globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_152_i2)) {
              break L0;
            }
            _152_i2 = (_152_i2).plus(_dafny.ONE);
            let _153_v10;
            let _nw27 = Array((new BigNumber(28)).toNumber());
            _nw27[(_dafny.ZERO).toNumber()] = _142_v3;
            _nw27[(_dafny.ONE).toNumber()] = _142_v3;
            _nw27[(new BigNumber(2)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(3)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(4)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(5)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(6)).toNumber()] = !(_142_v3);
            _nw27[(new BigNumber(7)).toNumber()] = _module.__default.fm0(_142_v3, _142_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(558)), ((_154_v0) => function (_155_i3) {
              return _154_v0;
            })(_138_v0)), _147_globalState);
            _nw27[(new BigNumber(8)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(9)).toNumber()] = false;
            _nw27[(new BigNumber(10)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(11)).toNumber()] = false;
            _nw27[(new BigNumber(12)).toNumber()] = false;
            _nw27[(new BigNumber(13)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(14)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(15)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(16)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(17)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(18)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(19)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(20)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(21)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(22)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(23)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(24)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(25)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(26)).toNumber()] = _142_v3;
            _nw27[(new BigNumber(27)).toNumber()] = _module.__default.fm0(_142_v3, true, _dafny.Seq.UnicodeFromString("w"), _147_globalState);
            _153_v10 = _nw27;
            let _156_v11;
            _156_v11 = _dafny.MultiSet.fromElements(_153_v10, _153_v10);
            _156_v11 = (((_156_v11).update(_153_v10, _module.__default.abs(_144_v5))).Difference(_dafny.MultiSet.fromElements(_153_v10))).Difference(_156_v11);
            let _index20 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_153_v10).length));
            (_153_v10)[_index20] = true;
            let _157_v12;
            _157_v12 = _dafny.MultiSet.fromElements((_151_v9)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), ((_158_v7) => function (_159_i4) {
              return new BigNumber((_158_v7).length);
            })(_146_v7))).length), new BigNumber((_151_v9).length))]);
            let _160_v13;
            _160_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_157_v12).Union(_157_v12)).cardinality()),_142_v3);
            let _index21 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_153_v10).length));
            let _rhs30 = _142_v3;
            let _rhs31 = _160_v13;
            let _lhs14 = _153_v10;
            let _lhs15 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_153_v10).length));
            _lhs14[_lhs15] = _rhs30;
            _160_v13 = _rhs31;
            let _161_v14;
            _161_v14 = _dafny.MultiSet.fromElements(_145_v6);
            let _162_v15;
            let _nw28 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _162_v15 = _nw28;
            let _163_v16;
            _163_v16 = _dafny.Map.Empty.slice().updateUnsafe(_161_v14,_162_v15);
            _163_v16 = (_163_v16).update(_161_v14, _162_v15);
            _142_v3 = ((_153_v10)[_module.__default.safeIndex(new BigNumber(42), new BigNumber((_153_v10).length))]) && ((_153_v10)[_module.__default.safeIndex(new BigNumber(42), new BigNumber((_153_v10).length))]);
          }
        }
      }
      let _164_v17;
      let _165_v18;
      let _out0;
      let _out1;
      let _outcollector0 = _module.__default.m0(_147_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _164_v17 = _out0;
      _165_v18 = _out1;
      let _hi1 = _164_v17;
      for (let _166_i5 = _module.__default.safeDivisionInt(_144_v5, _164_v17); _166_i5.isLessThan(_hi1); _166_i5 = _166_i5.plus(_dafny.ONE)) {
        _142_v3 = true;
        let _167_v19;
        _167_v19 = _dafny.Map.Empty.slice().updateUnsafe(_139_v1,new BigNumber((_145_v6).length));
        let _168_v20;
        _168_v20 = _dafny.Map.Empty.slice().updateUnsafe(_138_v0,_167_v19);
        _168_v20 = (_168_v20).update(_138_v0, _167_v19);
        _146_v7 = _146_v7;
        let _169_v21;
        let _170_v22;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m0(_147_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _169_v21 = _out2;
        _170_v22 = _out3;
      }
      (_147_globalState).f8 = _165_v18;
      _142_v3 = _142_v3;
      let _171_v23;
      let _nw29 = Array((new BigNumber(27)).toNumber()).fill([]);
      _171_v23 = _nw29;
      let _172_v24;
      let _nw30 = Array((new BigNumber(8)).toNumber()).fill(false);
      _172_v24 = _nw30;
      let _index22 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length));
      (_171_v23)[_index22] = ((_142_v3) ? (_172_v24) : (_172_v24));
      let _index23 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length));
      (_171_v23)[_index23] = _172_v24;
      let _173_v25;
      let _init5 = ((_174_v0) => function (_175_i6) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-562)), ((_176_v0) => function (_177_i7) {
          return _176_v0;
        })(_174_v0)), _dafny.Seq.UnicodeFromString("uueqdwt"));
      })(_138_v0);
      let _nw31 = Array((new BigNumber(15)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw31.length); _i0_5++) {
        _nw31[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _173_v25 = _nw31;
      let _index24 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length));
      (_173_v25)[_index24] = _151_v9;
      let _index25 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length));
      (_173_v25)[_index25] = _151_v9;
      (_147_globalState).f9 = _145_v6;
      let _178_i8;
      _178_i8 = _dafny.ZERO;
      L1: {
        while (!(!(_142_v3) || ((_164_v17).isEqualTo(_165_v18)))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_178_i8)) {
              break L1;
            }
            _178_i8 = (_178_i8).plus(_dafny.ONE);
            _138_v0 = new _dafny.CodePoint('h'.codePointAt(0));
            let _179_v26;
            _179_v26 = _dafny.MultiSet.fromElements(_143_v4);
            let _index26 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_139_v1).length));
            (_139_v1)[_index26] = _module.__default.fm1(_142_v3, new BigNumber((_179_v26).cardinality()), _165_v18, _147_globalState);
            let _180_v27;
            _180_v27 = _dafny.MultiSet.fromElements(_138_v0, _138_v0);
            let _index27 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_139_v1).length));
            (_139_v1)[_index27] = _module.__default.fm1((_180_v27).IsSubsetOf(_180_v27), ((_dafny.ZERO).minus(_164_v17)).multipliedBy(_165_v18), _164_v17, _147_globalState);
            let _181_v28;
            let _init6 = ((_182_v18) => function (_183_i9) {
              return (_183_i9).multipliedBy(_182_v18);
            })(_165_v18);
            let _nw32 = Array((new BigNumber(24)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw32.length); _i0_6++) {
              _nw32[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _181_v28 = _nw32;
            let _184_v29;
            _184_v29 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_181_v28);
            let _185_v30;
            _185_v30 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_142_v3),_181_v28);
            let _186_v31;
            _186_v31 = _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(_145_v6,_181_v28));
            let _187_v32;
            let _nw33 = Array((new BigNumber(22)).toNumber());
            _nw33[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_145_v6,(((_184_v29).contains(_142_v3)) ? ((_184_v29).get(_142_v3)) : (_181_v28)))).Merge(_185_v30);
            _nw33[(_dafny.ONE).toNumber()] = (_185_v30).Merge(_185_v30);
            _nw33[(new BigNumber(2)).toNumber()] = ((_142_v3) ? (_dafny.Map.Empty.slice().updateUnsafe(_145_v6,_181_v28)) : (_185_v30));
            _nw33[(new BigNumber(3)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(4)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(5)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(6)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(7)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(8)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(9)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(10)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(11)).toNumber()] = (_186_v31).dtor_cf0;
            _nw33[(new BigNumber(12)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(13)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(14)).toNumber()] = (_185_v30).Merge(_185_v30);
            _nw33[(new BigNumber(15)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(16)).toNumber()] = _185_v30;
            _nw33[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_145_v6,_181_v28);
            _nw33[(new BigNumber(18)).toNumber()] = ((_186_v31).dtor_cf0).Merge(_185_v30);
            _nw33[(new BigNumber(19)).toNumber()] = (_186_v31).dtor_cf0;
            _nw33[(new BigNumber(20)).toNumber()] = (_185_v30).update(_145_v6, _181_v28);
            _nw33[(new BigNumber(21)).toNumber()] = _185_v30;
            _187_v32 = _nw33;
            let _index28 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_187_v32).length));
            (_187_v32)[_index28] = _185_v30;
            let _188_v33;
            _188_v33 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_185_v30);
            let _index29 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_187_v32).length));
            (_187_v32)[_index29] = ((_142_v3) ? (_185_v30) : ((((_188_v33).contains(_142_v3)) ? ((_188_v33).get(_142_v3)) : (_185_v30))));
            (_147_globalState).f8 = _165_v18;
          }
        }
      }
      let _189_i10;
      _189_i10 = _dafny.ZERO;
      L2: {
        while (_dafny.Seq.IsPrefixOf((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], (_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))])) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_189_i10)) {
              break L2;
            }
            _189_i10 = (_189_i10).plus(_dafny.ONE);
            let _190_v34;
            let _nw34 = new _module.C0();
            _nw34.__ctor(false, _146_v7);
            _190_v34 = _nw34;
            let _index30 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_139_v1).length));
            (_139_v1)[_index30] = _144_v5;
            let _index31 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_139_v1).length));
            (_139_v1)[_index31] = new BigNumber(-698);
            let _191_v35;
            let _nw35 = Array((new BigNumber(29)).toNumber());
            _nw35[(_dafny.ZERO).toNumber()] = _150_v8;
            _nw35[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements(_190_v34.f19, _142_v3, _190_v34.f19, (((_146_v7).contains(_142_v3)) ? ((_146_v7).get(_142_v3)) : (_142_v3)))).Difference(_dafny.MultiSet.fromElements(_190_v34.f19, _142_v3));
            _nw35[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_module.__default.fm0(_142_v3, _142_v3, _module.__default.fm4(_190_v34.f19, _142_v3, !(_142_v3), new BigNumber((_143_v4).length), _147_globalState), _147_globalState), _190_v34.f19);
            _nw35[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_145_v6);
            _nw35[(new BigNumber(4)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(5)).toNumber()] = (_150_v8).update(_142_v3, _module.__default.abs((_139_v1)[_module.__default.safeIndex(new BigNumber(51), new BigNumber((_139_v1).length))]));
            _nw35[(new BigNumber(6)).toNumber()] = (_150_v8).Difference(_dafny.MultiSet.fromElements(_142_v3));
            _nw35[(new BigNumber(7)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(8)).toNumber()] = (_150_v8).update(_190_v34.f19, _module.__default.abs(_165_v18));
            _nw35[(new BigNumber(9)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(10)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_190_v34.f19));
            _nw35[(new BigNumber(12)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(13)).toNumber()] = (_150_v8).update(false, _module.__default.abs((_139_v1)[_module.__default.safeIndex(new BigNumber(51), new BigNumber((_139_v1).length))]));
            _nw35[(new BigNumber(14)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(15)).toNumber()] = (_150_v8).Difference(_dafny.MultiSet.fromElements(_190_v34.f19, _190_v34.f19));
            _nw35[(new BigNumber(16)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.fromElements(_190_v34.f19, _190_v34.f19, _190_v34.f19);
            _nw35[(new BigNumber(18)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(19)).toNumber()] = (_150_v8).Union(_dafny.MultiSet.fromElements(_142_v3));
            _nw35[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.fromElements(_142_v3);
            _nw35[(new BigNumber(21)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(22)).toNumber()] = (_150_v8).update(_142_v3, _module.__default.abs(_164_v17));
            _nw35[(new BigNumber(23)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(24)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(25)).toNumber()] = _dafny.MultiSet.fromElements(_190_v34.f19, _142_v3, _142_v3);
            _nw35[(new BigNumber(26)).toNumber()] = _150_v8;
            _nw35[(new BigNumber(27)).toNumber()] = (_150_v8).update(_module.__default.fm0(_142_v3, _142_v3, (_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _147_globalState), _module.__default.abs((_139_v1)[_module.__default.safeIndex(new BigNumber(51), new BigNumber((_139_v1).length))]));
            _nw35[(new BigNumber(28)).toNumber()] = _150_v8;
            _191_v35 = _nw35;
            _191_v35 = (_module.D1.create_DC3(_191_v35)).dtor_cf2;
            (_190_v34).f19 = _190_v34.f19;
          }
        }
      }
      let _source1 = _module.D0.create_DC1();
      if (_source1.is_DC1) {
        let _192_v36;
        _192_v36 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(!(_142_v3), _142_v3, (_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _147_globalState),_module.__default.fm5(_150_v8, _147_globalState));
        let _193_v37;
        _193_v37 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("odbnb"));
        let _194_v38;
        _194_v38 = _module.D1.create_DC4((((_192_v36).contains(!(_142_v3))) ? ((_192_v36).get(!(_142_v3))) : (_193_v37)));
        let _195_v39;
        _195_v39 = _dafny.Seq.of(new BigNumber(40));
        let _196_v40;
        _196_v40 = _dafny.Seq.of(_195_v39, _module.__default.fm7(_147_globalState));
        let _pat_let_tv2 = _193_v37;
        _194_v38 = function (_pat_let4_0) {
          return function (_197_dt__update__tmp_h0) {
            return function (_pat_let5_0) {
              return function (_198_dt__update_hcf3_h0) {
                return _module.D1.create_DC4(_198_dt__update_hcf3_h0);
              }(_pat_let5_0);
            }(_pat_let_tv2);
          }(_pat_let4_0);
        }(_module.__default.fm6(_196_v40, _142_v3, _142_v3, _147_globalState));
        _142_v3 = (_142_v3) && (_142_v3);
        let _199_v41;
        _199_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7(_147_globalState),_142_v3);
        _199_v41 = (_199_v41).update(_195_v39, (_142_v3) === (_142_v3));
        let _index32 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_172_v24).length));
        (_172_v24)[_index32] = _142_v3;
        let _index33 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_172_v24).length));
        (_172_v24)[_index33] = (new BigNumber((_143_v4).length)).isLessThanOrEqualTo(new BigNumber(((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]).length));
      } else if (_source1.is_DC0) {
        let _200___mcc_h0 = (_source1).cf0;
        let _201_cf0 = _200___mcc_h0;
        (_147_globalState).f8 = new BigNumber(551);
        let _202_v42;
        let _nw36 = new _module.C0();
        _nw36.__ctor(!(_142_v3) || (_142_v3), ((_146_v7)).Merge(_146_v7));
        _202_v42 = _nw36;
        _202_v42 = _202_v42;
        (_147_globalState).f8 = _144_v5;
        let _arr0 = (_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))];
        let _index34 = _module.__default.safeIndex(new BigNumber(386), new BigNumber(((_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))]).length));
        _arr0[_index34] = _142_v3;
        let _arr1 = (_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))];
        let _index35 = _module.__default.safeIndex(new BigNumber(386), new BigNumber(((_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))]).length));
        _arr1[_index35] = (_165_v18).isLessThan(new BigNumber(-638));
      } else {
        let _203___mcc_h1 = (_source1).cf1;
        let _204_cf1 = _203___mcc_h1;
        let _205_v43;
        let _nw37 = Array((new BigNumber(21)).toNumber());
        _205_v43 = _nw37;
        _205_v43 = _205_v43;
        if (_142_v3) {
          (_147_globalState).f7 = _144_v5;
          let _index36 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_172_v24).length));
          (_172_v24)[_index36] = _142_v3;
          let _206_v44;
          _206_v44 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,new _dafny.CodePoint('g'.codePointAt(0)));
          let _207_v45;
          _207_v45 = _dafny.Set.fromElements(_144_v5, _144_v5, _module.__default.fm1(_142_v3, _144_v5, new BigNumber((_206_v44).length), _147_globalState), _module.__default.safeDivisionInt(new BigNumber(-178), _164_v17), _165_v18);
          let _index37 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_172_v24).length));
          let _rhs32 = ((_module.__default.fm0(_142_v3, _142_v3, _151_v9, _147_globalState)) ? (_144_v5) : (_165_v18));
          let _rhs33 = _dafny.Seq.update((((_164_v17).isLessThanOrEqualTo(_165_v18)) ? (_145_v6) : (_module.__default.fm8(_165_v18, _142_v3, _147_globalState))), _module.__default.safeIndex(_144_v5, new BigNumber(((((_164_v17).isLessThanOrEqualTo(_165_v18)) ? (_145_v6) : (_module.__default.fm8(_165_v18, _142_v3, _147_globalState)))).length)), !(_142_v3) || (!(false)));
          let _rhs34 = ((_142_v3) ? (!(_dafny.Seq.IsPrefixOf((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], (_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]))) : ((new BigNumber((_dafny.MultiSet.fromElements((_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))], _172_v24, _172_v24)).cardinality())).isLessThan(new BigNumber(-337))));
          let _rhs35 = _207_v45;
          let _lhs16 = _147_globalState;
          let _lhs17 = _147_globalState;
          let _lhs18 = _172_v24;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_172_v24).length));
          _lhs16.f8 = _rhs32;
          _lhs17.f9 = _rhs33;
          _lhs18[_lhs19] = _rhs34;
          _207_v45 = _rhs35;
          let _208_v46;
          let _209_v47;
          let _out4;
          let _out5;
          let _outcollector2 = _module.__default.m0(_147_globalState);
          _out4 = _outcollector2[0];
          _out5 = _outcollector2[1];
          _208_v46 = _out4;
          _209_v47 = _out5;
          let _210_v48;
          let _nw38 = new _module.C0();
          _nw38.__ctor(!(true), _146_v7);
          _210_v48 = _nw38;
          _210_v48 = _210_v48;
          _142_v3 = (_144_v5).isLessThanOrEqualTo((_210_v48).fm2(_138_v0, _module.__default.fm0(_210_v48.f19, (_172_v24)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_172_v24).length))], _151_v9, _147_globalState), _147_globalState));
        } else {
          let _211_v49;
          let _212_v50;
          let _out6;
          let _out7;
          let _outcollector3 = _module.__default.m0(_147_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _211_v49 = _out6;
          _212_v50 = _out7;
          let _index38 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_172_v24).length));
          (_172_v24)[_index38] = _dafny.Seq.IsProperPrefixOf(_module.__default.fm8(_164_v17, _142_v3, _147_globalState), _module.__default.fm8(_211_v49, _142_v3, _147_globalState));
          let _index39 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_172_v24).length));
          (_172_v24)[_index39] = (!(_142_v3)) && (_142_v3);
          let _213_v51;
          _213_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mnvbm"),_212_v50);
          (_147_globalState).f7 = (((_213_v51).contains(_151_v9)) ? ((_213_v51).get(_151_v9)) : (_212_v50));
          let _214_v52;
          _214_v52 = _dafny.Map.Empty.slice().updateUnsafe(_211_v49,_211_v49);
          let _215_v53;
          let _nw39 = new _module.C0();
          _nw39.__ctor(_142_v3, _module.__default.fm9((((_214_v52).contains((new BigNumber(((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]).length)))) ? ((_214_v52).get((new BigNumber(((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]).length)))) : (_211_v49)), _144_v5, new BigNumber((_151_v9).length), _147_globalState));
          _215_v53 = _nw39;
          let _216_v54;
          _216_v54 = _dafny.Map.Empty.slice().updateUnsafe(_139_v1,(_212_v50).minus(_144_v5));
          _216_v54 = (_216_v54).update(_139_v1, (_dafny.ZERO).minus(_211_v49));
        }
        _138_v0 = _138_v0;
        _142_v3 = !(true);
      }
      let _217_i11;
      _217_i11 = _dafny.ZERO;
      L3: {
        while (_142_v3) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_217_i11)) {
              break L3;
            }
            _217_i11 = (_217_i11).plus(_dafny.ONE);
            let _rhs36 = !(_142_v3) || (_142_v3);
            let _rhs37 = _164_v17;
            let _rhs38 = new BigNumber(744);
            let _lhs20 = _147_globalState;
            let _lhs21 = _147_globalState;
            _142_v3 = _rhs36;
            _lhs20.f8 = _rhs37;
            _lhs21.f8 = _rhs38;
            if (_module.__default.fm0(_142_v3, _142_v3, _dafny.Seq.UnicodeFromString("xahmwdx"), _147_globalState)) {
              let _218_v55;
              _218_v55 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_151_v9).length)).multipliedBy(_165_v18),new _dafny.CodePoint('p'.codePointAt(0)));
              _218_v55 = (_218_v55).update(_144_v5, _138_v0);
              let _219_v56;
              let _nw40 = new _module.C0();
              _nw40.__ctor(_142_v3, _146_v7);
              _219_v56 = _nw40;
              let _220_v57;
              let _221_v58;
              let _out8;
              let _out9;
              let _outcollector4 = _module.__default.m0(_147_globalState);
              _out8 = _outcollector4[0];
              _out9 = _outcollector4[1];
              _220_v57 = _out8;
              _221_v58 = _out9;
              let _222_v59;
              _222_v59 = !(_142_v3);
              (_219_v56).f19 = (_222_v59);
              let _223_v60;
              _223_v60 = _dafny.Seq.of((((_150_v8).contains(_219_v56.f19)) ? ((_150_v8).get(_219_v56.f19)) : (_165_v18)), _144_v5);
              let _224_v61;
              _224_v61 = _dafny.Map.Empty.slice().updateUnsafe(_219_v56,_223_v60);
              _224_v61 = (_224_v61).update(_219_v56, _223_v60);
            } else {
              let _225_v62;
              let _nw41 = new _module.C0();
              _nw41.__ctor(_142_v3, (_146_v7).Merge(_146_v7));
              _225_v62 = _nw41;
              _146_v7 = (_146_v7).update(_module.__default.fm0(_142_v3, _142_v3, (_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _147_globalState), !(_142_v3));
              let _226_v63;
              _226_v63 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_module.__default.fm1(_142_v3, _165_v18, (_dafny.ZERO).minus(_144_v5), _147_globalState), _164_v17),_module.__default.safeDivisionInt(_164_v17, _165_v18));
              _226_v63 = (_226_v63).update((_165_v18).multipliedBy(_165_v18), new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_164_v17), (_225_v62).fm2(_138_v0, _225_v62.f19, _147_globalState), _144_v5)).cardinality()));
              _146_v7 = (_146_v7).update(false, _142_v3);
              let _227_v64;
              let _nw42 = new _module.C0();
              _nw42.__ctor((_225_v62.f19) === ((_145_v6)[_module.__default.safeIndex(_165_v18, new BigNumber((_145_v6).length))]), _146_v7);
              _227_v64 = _nw42;
            }
            let _index40 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_139_v1).length));
            (_139_v1)[_index40] = _165_v18;
            let _index41 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_139_v1).length));
            let _rhs39 = new BigNumber(((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]).length);
            let _rhs40 = _module.__default.safeDivisionInt(_164_v17, _164_v17);
            let _rhs41 = _144_v5;
            let _lhs22 = _147_globalState;
            let _lhs23 = _147_globalState;
            let _lhs24 = _139_v1;
            let _lhs25 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_139_v1).length));
            _lhs22.f8 = _rhs39;
            _lhs23.f7 = _rhs40;
            _lhs24[_lhs25] = _rhs41;
            let _228_v65;
            _228_v65 = _dafny.MultiSet.fromElements(_dafny.Seq.update((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _module.__default.safeIndex(_164_v17, new BigNumber(((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]).length)), _138_v0), _151_v9);
            let _229_v66;
            _229_v66 = _dafny.Seq.of(_dafny.Seq.update((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _module.__default.safeIndex(_144_v5, new BigNumber(((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]).length)), _138_v0));
            _142_v3 = (_dafny.MultiSet.FromArray(_229_v66)).IsProperSubsetOf((_228_v65).Union(_228_v65));
          }
        }
      }
      _142_v3 = _142_v3;
      if (_142_v3) {
        if ((false) && (_142_v3)) {
          let _230_v67;
          let _nw43 = new _module.C0();
          _nw43.__ctor(_142_v3, _146_v7);
          _230_v67 = _nw43;
          _230_v67 = _230_v67;
          let _init7 = ((_231_v5) => function (_232_i12) {
            return (_232_i12).multipliedBy(_231_v5);
          })(_144_v5);
          let _nw44 = Array((new BigNumber(15)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw44.length); _i0_7++) {
            _nw44[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _139_v1 = _nw44;
          (_147_globalState).f8 = _165_v18;
          let _arr2 = (_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))];
          let _index42 = _module.__default.safeIndex(new BigNumber(41), new BigNumber(((_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))]).length));
          _arr2[_index42] = !(_144_v5).isEqualTo(_164_v17);
          let _233_v68;
          _233_v68 = _dafny.Set.fromElements(_164_v17, new BigNumber(151));
          let _234_v70;
          _234_v70 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll3 = new _dafny.Set();
            for (const _compr_3 of _dafny.IntegerRange(new BigNumber(168), new BigNumber(106))) {
              let _235_v69 = _compr_3;
              if (((new BigNumber(168)).isLessThanOrEqualTo(_235_v69)) && ((_235_v69).isLessThan(new BigNumber(106)))) {
                _coll3.add((_235_v69).multipliedBy(_144_v5));
              }
            }
            return _coll3;
          }(),_145_v6);
          let _arr3 = (_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))];
          let _index43 = _module.__default.safeIndex(new BigNumber(41), new BigNumber(((_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))]).length));
          _arr3[_index43] = !(_234_v70).contains(_233_v68);
          _230_v67 = _230_v67;
        } else {
          let _236_v71;
          let _nw45 = new _module.C0();
          _nw45.__ctor((((_146_v7).contains(_142_v3)) ? ((_146_v7).get(_142_v3)) : (_142_v3)), _146_v7);
          _236_v71 = _nw45;
          let _237_v72;
          _237_v72 = _module.D5.create_DC9(_236_v71);
          let _pat_let_tv3 = _236_v71;
          _236_v71 = (function (_pat_let6_0) {
            return function (_238_dt__update__tmp_h1) {
              return function (_pat_let7_0) {
                return function (_239_dt__update_hcf8_h0) {
                  return _module.D5.create_DC9(_239_dt__update_hcf8_h0);
                }(_pat_let7_0);
              }(_pat_let_tv3);
            }(_pat_let6_0);
          }(_237_v72)).dtor_cf8;
          let _240_v73;
          let _nw46 = Array((new BigNumber(8)).toNumber());
          _nw46[(_dafny.ZERO).toNumber()] = _150_v8;
          _nw46[(_dafny.ONE).toNumber()] = _150_v8;
          _nw46[(new BigNumber(2)).toNumber()] = _150_v8;
          _nw46[(new BigNumber(3)).toNumber()] = _150_v8;
          _nw46[(new BigNumber(4)).toNumber()] = _150_v8;
          _nw46[(new BigNumber(5)).toNumber()] = _150_v8;
          _nw46[(new BigNumber(6)).toNumber()] = _150_v8;
          _nw46[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(false);
          _240_v73 = _nw46;
          let _nw47 = Array((new BigNumber(19)).toNumber()).fill(_dafny.MultiSet.Empty);
          _240_v73 = _nw47;
          let _241_v74;
          _241_v74 = _dafny.Seq.of((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _151_v9);
          _151_v9 = _dafny.Seq.Concat((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _dafny.Seq.Concat((_241_v74)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_241_v74).length))], _151_v9));
          _236_v71 = _236_v71;
          let _index44 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_139_v1).length));
          (_139_v1)[_index44] = new BigNumber(-834);
          let _242_v75;
          _242_v75 = _dafny.Set.fromElements((_236_v71).fm2(_138_v0, !(!(_236_v71.f19)), _147_globalState), _165_v18);
          let _243_v76;
          _243_v76 = _dafny.MultiSet.fromElements(new BigNumber((_150_v8).cardinality()));
          let _244_v77;
          _244_v77 = _dafny.Map.Empty.slice().updateUnsafe((_243_v76).update(new BigNumber((_151_v9).length), _module.__default.abs(_144_v5)),_165_v18);
          let _245_v78;
          _245_v78 = _dafny.Seq.of(_164_v17);
          let _index45 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_139_v1).length));
          let _rhs42 = (_245_v78)[_module.__default.safeIndex((_dafny.ZERO).minus(_165_v18), new BigNumber((_245_v78).length))];
          let _rhs43 = ((_142_v3) ? (_242_v75) : (_242_v75));
          let _rhs44 = _module.__default.fm10(_245_v78, _module.__default.fm11(_138_v0, _236_v71.f19, _164_v17, _165_v18, _147_globalState), _147_globalState);
          let _rhs45 = (_244_v77).Merge(_244_v77);
          let _rhs46 = _142_v3;
          let _lhs26 = _139_v1;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_139_v1).length));
          _lhs26[_lhs27] = _rhs42;
          _242_v75 = _rhs43;
          _138_v0 = _rhs44;
          _244_v77 = _rhs45;
          _142_v3 = _rhs46;
        }
        let _246_v79;
        _246_v79 = _dafny.Seq.of(_164_v17);
        let _247_v80;
        _247_v80 = _dafny.Seq.of(_246_v79, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-204)), ((_248_v17) => function (_249_i13) {
          return new BigNumber((_dafny.MultiSet.fromElements(_248_v17, new BigNumber(362))).cardinality());
        })(_164_v17)));
        _143_v4 = _module.__default.fm12(_247_v80, (_144_v5).isLessThanOrEqualTo(_144_v5), (((_146_v7).contains(!(false))) ? ((_146_v7).get(!(false))) : (_142_v3)), _147_globalState);
        if (!(_142_v3) || ((((_145_v6)[_module.__default.safeIndex(_165_v18, new BigNumber((_145_v6).length))]) ? (_142_v3) : (_142_v3)))) {
          let _250_v81;
          let _nw48 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Set.Empty);
          _250_v81 = _nw48;
          _250_v81 = _250_v81;
          _171_v23 = _171_v23;
          let _251_v82;
          _251_v82 = _dafny.Seq.of(_172_v24, _172_v24, (_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))]);
          _251_v82 = _dafny.Seq.Concat(_251_v82, _251_v82);
          let _index46 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_172_v24).length));
          (_172_v24)[_index46] = !(true);
          let _index47 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_172_v24).length));
          (_172_v24)[_index47] = ((((_146_v7).contains(_142_v3)) ? ((_146_v7).get(_142_v3)) : (_142_v3))) && (_142_v3);
          (_147_globalState).f8 = _164_v17;
        } else {
          let _252_v83;
          let _nw49 = new _module.C0();
          _nw49.__ctor(_142_v3, _146_v7);
          _252_v83 = _nw49;
          _252_v83 = _252_v83;
          (_147_globalState).f7 = _module.__default.safeModuloInt(_144_v5, _165_v18);
          let _253_v84;
          _253_v84 = _dafny.Seq.of(_252_v83);
          _252_v83 = (_253_v84)[_module.__default.safeIndex((_144_v5).minus(_144_v5), new BigNumber((_253_v84).length))];
          (_147_globalState).f8 = new BigNumber((_dafny.Seq.of(_252_v83.f19, !(_142_v3), _142_v3, ((_142_v3) ? (_252_v83.f19) : (_142_v3)))).length);
          let _index48 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_172_v24).length));
          (_172_v24)[_index48] = _252_v83.f19;
          let _254_v85;
          _254_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(442),_142_v3);
          let _index49 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_172_v24).length));
          let _rhs47 = _252_v83.f19;
          let _rhs48 = _254_v85;
          let _lhs28 = _172_v24;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_172_v24).length));
          _lhs28[_lhs29] = _rhs47;
          _254_v85 = _rhs48;
        }
        let _255_v87;
        _255_v87 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_146_v7).length)), new BigNumber(693), _164_v17);
        let _256_v88;
        _256_v88 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(_164_v17), _255_v87, _255_v87, _255_v87, _255_v87);
        let _257_v89;
        _257_v89 = _module.D5.create_DC10(_dafny.Set.fromElements(_144_v5, _144_v5, _165_v18), function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of (_256_v88).Elements) {
    let _258_v86 = _compr_4;
    if ((_256_v88).contains(_258_v86)) {
      _coll4.push([_258_v86,_151_v9]);
    }
  }
  return _coll4;
}());
        let _259_v90;
        _259_v90 = _dafny.Seq.of(((_142_v3) ? (_257_v89) : (_257_v89)), _257_v89, _257_v89);
        let _260_v91;
        _260_v91 = _dafny.Set.fromElements(_144_v5);
        let _261_v92;
        _261_v92 = _dafny.Map.Empty.slice().updateUnsafe(_255_v87,(_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]);
        let _rhs49 = _139_v1;
        let _rhs50 = _dafny.Seq.update(_dafny.Seq.of(_257_v89, _module.D5.create_DC10(_260_v91, _261_v92)), _module.__default.safeIndex(_164_v17, new BigNumber((_dafny.Seq.of(_257_v89, _module.D5.create_DC10(_260_v91, _261_v92))).length)), _257_v89);
        let _rhs51 = _138_v0;
        _139_v1 = _rhs49;
        _259_v90 = _rhs50;
        _138_v0 = _rhs51;
        let _262_v93;
        _262_v93 = _151_v9;
        let _263_v94;
        let _nw50 = new _module.C0();
        _nw50.__ctor(_module.__default.fm0(_142_v3, _142_v3, (_262_v93), _147_globalState), _146_v7);
        _263_v94 = _nw50;
      } else {
        if ((_144_v5).isLessThan((_164_v17).multipliedBy(new BigNumber(821)))) {
          let _index50 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length));
          (_139_v1)[_index50] = (_165_v18).minus(_module.__default.fm1(_142_v3, _165_v18, (_dafny.ZERO).minus(_164_v17), _147_globalState));
          let _264_v95;
          _264_v95 = _dafny.Seq.of(_139_v1);
          let _265_v96;
          _265_v96 = _dafny.Map.Empty.slice().updateUnsafe(_165_v18,_165_v18);
          let _index51 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length));
          let _rhs52 = ((_142_v3) ? (new BigNumber((_151_v9).length)) : (((_dafny.ZERO).minus(_164_v17)).minus(_164_v17)));
          let _rhs53 = (_264_v95)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_142_v3, _142_v3)).length), new BigNumber((_264_v95).length))];
          let _rhs54 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_165_v18), _module.__default.safeModuloInt((((_265_v96).contains(_144_v5)) ? ((_265_v96).get(_144_v5)) : (_164_v17)), _144_v5));
          let _lhs30 = _147_globalState;
          let _lhs31 = _139_v1;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length));
          _lhs30.f7 = _rhs52;
          _139_v1 = _rhs53;
          _lhs31[_lhs32] = _rhs54;
          let _266_v97;
          _266_v97 = _dafny.MultiSet.fromElements(_165_v18, (_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))]);
          let _267_v98;
          _267_v98 = _dafny.Map.Empty.slice().updateUnsafe(_266_v97,new BigNumber(((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))]).length));
          let _268_v99;
          _268_v99 = _dafny.Map.Empty.slice().updateUnsafe((_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))],_267_v98);
          let _269_v100;
          _269_v100 = _dafny.Seq.of(_146_v7, _146_v7, _146_v7);
          let _270_v102;
          let _nw51 = new _module.C0();
          _nw51.__ctor(_142_v3, _146_v7);
          _270_v102 = _nw51;
          let _271_v103;
          _271_v103 = _dafny.Map.Empty.slice().updateUnsafe(_270_v102,_142_v3);
          let _272_v104;
          let _nw52 = Array((new BigNumber(8)).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = ((_142_v3) ? ((_267_v98).update((_266_v97).update((_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))], _module.__default.abs(_165_v18)), (_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))])) : ((((_268_v99).contains(new BigNumber(((_269_v100)[_module.__default.safeIndex(_165_v18, new BigNumber((_269_v100).length))]).length))) ? ((_268_v99).get(new BigNumber(((_269_v100)[_module.__default.safeIndex(_165_v18, new BigNumber((_269_v100).length))]).length))) : (_267_v98))));
          _nw52[(_dafny.ONE).toNumber()] = (function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of (_dafny.Seq.of(_266_v97)).Elements) {
              let _273_v101 = _compr_5;
              if (_dafny.Seq.contains(_dafny.Seq.of(_266_v97), _273_v101)) {
                _coll5.push([_273_v101,_144_v5]);
              }
            }
            return _coll5;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_266_v97,_144_v5));
          _nw52[(new BigNumber(2)).toNumber()] = _267_v98;
          _nw52[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_266_v97,_144_v5)).update(_266_v97, new BigNumber((_271_v103).length));
          _nw52[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_164_v17, new BigNumber(274), (_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))], _164_v17),_164_v17);
          _nw52[(new BigNumber(5)).toNumber()] = _module.__default.fm13(_142_v3, (_dafny.ZERO).minus(_144_v5), _147_globalState);
          _nw52[(new BigNumber(6)).toNumber()] = _267_v98;
          _nw52[(new BigNumber(7)).toNumber()] = (_267_v98).Merge(_267_v98);
          _272_v104 = _nw52;
          let _index52 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_272_v104).length));
          (_272_v104)[_index52] = _267_v98;
          let _index53 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_272_v104).length));
          let _rhs55 = (_dafny.ZERO).minus((_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))]);
          let _rhs56 = _module.__default.fm1(true, _164_v17, (_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))], _147_globalState);
          let _rhs57 = _dafny.Seq.Concat(_dafny.Seq.Concat((_173_v25)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), ((_274_v0) => function (_275_i14) {
            return _274_v0;
          })(_138_v0))), _dafny.Seq.Concat(_151_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), ((_276_v0) => function (_277_i15) {
            return _276_v0;
          })(_138_v0))));
          let _rhs58 = _267_v98;
          let _lhs33 = _147_globalState;
          let _lhs34 = _272_v104;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_272_v104).length));
          _lhs33.f8 = _rhs55;
          _165_v18 = _rhs56;
          _151_v9 = _rhs57;
          _lhs34[_lhs35] = _rhs58;
          let _278_v105;
          _278_v105 = _dafny.Seq.of(_172_v24, _172_v24, _172_v24, _172_v24, (_171_v23)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_171_v23).length))]);
          _278_v105 = _278_v105;
          let _279_v106;
          _279_v106 = _dafny.Seq.of((_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))], (_dafny.ZERO).minus(_164_v17));
          let _280_v107;
          _280_v107 = _dafny.Seq.of(_279_v106, _279_v106, _279_v106, _dafny.Seq.of(_164_v17, new BigNumber((_dafny.Seq.of((_139_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_139_v1).length))])).length), _144_v5));
          let _281_v108;
          _281_v108 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("rjjbn"));
          let _index54 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length));
          let _rhs59 = _142_v3;
          let _rhs60 = _dafny.Seq.Concat((_280_v107)[_module.__default.safeIndex(_144_v5, new BigNumber((_280_v107).length))], _dafny.Seq.of(new BigNumber((_281_v108).length)));
          let _rhs61 = _142_v3;
          let _rhs62 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mwwcq"), _151_v9);
          let _lhs36 = _270_v102;
          let _lhs37 = _173_v25;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_173_v25).length));
          _142_v3 = _rhs59;
          _279_v106 = _rhs60;
          _lhs36.f19 = _rhs61;
          _lhs37[_lhs38] = _rhs62;
          (_270_v102).f19 = (false) === (_270_v102.f19);
        } else {
          _142_v3 = _142_v3;
          let _282_v109;
          let _nw53 = new _module.C0();
          _nw53.__ctor(_142_v3, _146_v7);
          _282_v109 = _nw53;
          let _index55 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_139_v1).length));
          (_139_v1)[_index55] = (_165_v18).multipliedBy(_164_v17);
          let _283_v110;
          _283_v110 = _dafny.Map.Empty.slice().updateUnsafe(_150_v8,_164_v17);
          let _index56 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_139_v1).length));
          (_139_v1)[_index56] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_283_v110).length), _module.__default.safeModuloInt(_144_v5, _165_v18)));
          (_282_v109).f19 = _282_v109.f19;
          (_147_globalState).f7 = ((_dafny.ZERO).minus((_139_v1)[_module.__default.safeIndex(new BigNumber(32), new BigNumber((_139_v1).length))])).minus(_164_v17);
        }
        let _284_v111;
        let _nw54 = new _module.C0();
        _nw54.__ctor((_145_v6)[_module.__default.safeIndex(_164_v17, new BigNumber((_145_v6).length))], _146_v7);
        _284_v111 = _nw54;
        _146_v7 = (_146_v7).update(_142_v3, _284_v111.f19);
        let _285_v112;
        _285_v112 = _dafny.Map.Empty.slice().updateUnsafe(_138_v0,_284_v111);
        _285_v112 = ((_142_v3) ? (((_285_v112).update(_138_v0, _284_v111)).update(new _dafny.CodePoint('w'.codePointAt(0)), _284_v111)) : (_285_v112));
        _142_v3 = _284_v111.f19;
      }
      _146_v7 = (_146_v7).update(_142_v3, _142_v3);
      process.stdout.write(_dafny.toString(_138_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_141_v2).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_142_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v4).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_144_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_145_v6, _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_147_globalState).f1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_147_globalState).f2).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f3).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState.f4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))),new BigNumber(172)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_147_globalState.f9, _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f13).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_150_v8).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_151_v9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_152_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_164_v17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_165_v18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_171_v23)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v24)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_173_v25)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_178_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_189_i10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_217_i11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf1) {
      let $dt = new D0(2);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf1) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC4(cf3) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC3(cf2) {
      let $dt = new D1(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC5(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.Set.Empty);
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
    static create_DC6(cf5) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
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
    static create_DC7(cf6) {
      let $dt = new D3(0);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.ZERO;
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
    static create_DC8(cf7) {
      let $dt = new D4(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
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
    static create_DC10(cf9, cf10) {
      let $dt = new D5(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC11(cf11, cf12, cf13, cf14, cf15) {
      let $dt = new D5(1);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC9(cf8) {
      let $dt = new D5(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC12(cf16) {
      let $dt = new D5(3);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC9" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf8 === other.cf8;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC10(_dafny.Set.Empty, _dafny.Map.Empty);
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
    static create_DC13(cf17) {
      let $dt = new D6(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + this.cf17.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
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
    static create_DC15(cf19, cf20, cf21, cf22) {
      let $dt = new D7(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC16(cf23, cf24, cf25, cf26, cf27) {
      let $dt = new D7(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC17(cf28, cf29, cf30, cf31) {
      let $dt = new D7(2);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC14(cf18) {
      let $dt = new D7(3);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30 && this.cf31 === other.cf31;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15(null, _dafny.ZERO, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f4 = _dafny.Map.Empty;
      this.f7 = _dafny.ZERO;
      this.f8 = _dafny.ZERO;
      this.f9 = _dafny.Seq.of();
      this._f0 = false;
      this._f1 = _dafny.Seq.UnicodeFromString("");
      this._f2 = _dafny.Map.Empty;
      this._f3 = _dafny.Set.Empty;
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f10 = false;
      this._f11 = [];
      this._f12 = false;
      this._f13 = _dafny.Map.Empty;
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
      this._f16 = false;
      this._f17 = false;
      this._f18 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f19 = false;
      this._f20 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f19, f20) {
      let _this = this;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber(((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length), new BigNumber(((_this).f20).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length))).length))).length))).Intersect((_dafny.Set.fromElements(new BigNumber(689))).Union(_dafny.Set.fromElements(new BigNumber(468), new BigNumber((_dafny.Seq.UnicodeFromString("doisvt")).length))))).length));
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      let _source2 = _module.D0.create_DC1();
      if (_source2.is_DC1) {
        return _module.D0.create_DC1();
      } else if (_source2.is_DC0) {
        let _286___mcc_h0 = (_source2).cf0;
        let _287_cf0 = _286___mcc_h0;
        return _module.D0.create_DC1();
      } else {
        let _288___mcc_h1 = (_source2).cf1;
        let _289_cf1 = _288___mcc_h1;
        return _module.D0.create_DC1();
      }
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
